const holder = artifacts.require("./WriteRequestHolder.sol");
const Web3 = require("web3");
const ganache = require('ganache-cli');
const web3 = new Web3(ganache.provider());

module.exports = (deployer) => {
    deployRRHolder(deployer);
};

const deployRRHolder =  async (deployer) => {
    deployer.deploy(holder);
}