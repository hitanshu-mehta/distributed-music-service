pragma solidity 0.5.5;

import "./openzeppelin/contracts/token/ERC20/ERC20Mintable.sol";
import "./openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";
import "./openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MusicToken is ERC20, ERC20Mintable, ERC20Detailed {
     constructor()
        ERC20Detailed("Music Token", "MTK", 18)
        public {

    }
}
