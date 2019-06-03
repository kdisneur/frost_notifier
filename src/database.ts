import { DynamoDB } from 'aws-sdk'
import { DocumentClient } from 'aws-sdk/clients/dynamodb'
import { ILogger } from './logger'

export interface IConfig {
  tableName: string
  logger: ILogger
}

const dynamoDB = new DynamoDB.DocumentClient()

interface IDatabase {
  hasItem: (config: IConfig, timestamp: Date) => Promise<boolean>
  putItem: (config: IConfig, timestamp: Date, value: boolean) => Promise<void>
}

const toSendingDate = (date: Date): string => {
  return date.toISOString().substring(0,10)
}

const hasItem = (config: IConfig, date: Date): Promise<boolean> => {
  const params = {
    TableName: config.tableName,
    Key: { 'SendingDate': toSendingDate(date) }
  }

  return new Promise((resolve, reject) => {
    dynamoDB.get(params, (error: Error, value: DocumentClient.GetItemOutput) => {
      if (error) {
        config.logger.error(`dynamodb query failed: ${error}`)
        return resolve(false)
      }

      return resolve(value.Item && value.Item['Done'])
    })
  })
}

const putItem = (config: IConfig, date: Date, done: boolean): Promise<void> => {
  const params = {
    TableName: config.tableName,
    Item: { SendingDate: toSendingDate(date), Done: done }
  }

  return new Promise((resolve, reject) => {
    dynamoDB.put(params, (error: Error, result): void => {
      if (error) {
        config.logger.error(`dynamodb command failed: ${error}`)
        return reject(error)
      }

      return resolve()
    })
  })
}

const database: IDatabase = { hasItem, putItem }

export default database
