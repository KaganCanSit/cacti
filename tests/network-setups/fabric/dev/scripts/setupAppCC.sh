CC_NAME=${1:-"simplestate"}

CHAINCODE_PATH=$PWD/../shared/chaincode
APP_CC_PATH=$PWD/../../../../samples/${CC_NAME}

echo "Copying ${CC_NAME}..."

##################### Copying app chaincode ###############
if [ -d "${CHAINCODE_PATH}/${CC_NAME}" ]; then
    echo "Deleting existing ${CC_NAME} chaincode folder to copy the latest..."
    rm -rf ${CHAINCODE_PATH}/${CC_NAME}
fi
cp -r ${APP_CC_PATH} ${CHAINCODE_PATH}/

echo "Done."
