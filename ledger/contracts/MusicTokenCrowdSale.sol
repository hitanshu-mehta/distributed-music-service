pragma solidity 0.5.5;
pragma experimental ABIEncoderV2;

import "./openzeppelin/contracts/crowdsale/Crowdsale.sol";
import "./openzeppelin/contracts/crowdsale/emission/MintedCrowdsale.sol";
import "./openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./openzeppelin/contracts/token/ERC20/ERC20Mintable.sol";
import "./MusicToken.sol";

contract MusicTokenCrowdSale is Crowdsale, MintedCrowdsale {
    constructor(uint256 _rate, address payable _wallet, IERC20 _token)
        MintedCrowdsale() 
        Crowdsale(_rate, _wallet, _token)
        public
    {
    }
}

contract MusicTokenCrowdsaleDeployer {
    constructor()
        public
    {
        // create a mintable token
        ERC20Mintable token = new MusicToken();

        // create the crowdsale and tell it about the token
        Crowdsale crowdsale = new MusicTokenCrowdSale(
            1,               // rate, still in TKNbits
            msg.sender,      // send Ether to the deployer
            token            // the token
        );
        // transfer the minter role from this contract (the default)
        // to the crowdsale, so it can mint tokens
        token.addMinter(address(crowdsale));
        token.renounceMinter();
    }
}