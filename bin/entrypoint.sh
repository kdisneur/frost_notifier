#! /usr/bin/env bash

COUNTRY=${COUNTRY:?"environment variable is missing"};
LANGUAGE=${LANGUAGE:?"environment variable is missing"};
POSTCODE=${POSTCODE:?"environment variable is missing"};
PHONE=${PHONE:?"environment variable is missing"};
OPENWEATHER_APIKEY=${OPENWEATHER_APIKEY:?"environment variable is missing"};
TWILIO_ACCOUNTSID=${TWILIO_ACCOUNTSID:?"environment variable is missing"};
TWILIO_TOKEN=${TWILIO_TOKEN:?"environment variable is missing"};
TWILIO_VIRTUALPHONENUMBER=${TWILIO_VIRTUALPHONENUMBER:?"environment variable is missing"};

CREDENTIALS_PATH=${CREDENTIALS_PATH:-"./credentials.json"}

BINARY=${BINARY:-./frost-notifier};

${BINARY} -f ${CREDENTIALS_PATH} init <<EOF
${OPENWEATHER_APIKEY}
${TWILIO_ACCOUNTSID}
${TWILIO_TOKEN}
${TWILIO_VIRTUALPHONENUMBER}

EOF

${BINARY} -f ${CREDENTIALS_PATH} -l ${LANGUAGE} ${COUNTRY} ${POSTCODE} ${PHONE}
