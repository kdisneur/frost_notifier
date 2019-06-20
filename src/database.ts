import { DynamoDB } from "aws-sdk";
import { DocumentClient } from "aws-sdk/clients/dynamodb";
import * as logger from "./logger";

export interface Config {
  tableName: string;
  logger: logger.Logger;
}

interface Database {
  hasItem: (config: Config, timestamp: Date) => Promise<boolean>;
  putItem: (config: Config, timestamp: Date, value: boolean) => Promise<void>;
}

const dynamoDB = new DynamoDB.DocumentClient();

const toSendingDate = (date: Date): string => {
  return date.toISOString().substring(0, 10);
};

export const hasItem = (config: Config, date: Date): Promise<boolean> => {
  const params = {
    TableName: config.tableName,
    Key: { SendingDate: toSendingDate(date) }
  };

  return new Promise(resolve => {
    dynamoDB.get(
      params,
      (error: Error, value: DocumentClient.GetItemOutput) => {
        if (error) {
          config.logger.error(`dynamodb query failed: ${error}`);
          return resolve(false);
        }

        return resolve(value.Item && value.Item["Done"]);
      }
    );
  });
};

export const putItem = (
  config: Config,
  date: Date,
  done: boolean
): Promise<void> => {
  const params = {
    TableName: config.tableName,
    Item: { SendingDate: toSendingDate(date), Done: done }
  };

  return new Promise((resolve, reject) => {
    dynamoDB.put(params, (error: Error): void => {
      if (error) {
        config.logger.error(`dynamodb command failed: ${error}`);
        return reject(error);
      }

      return resolve();
    });
  });
};
