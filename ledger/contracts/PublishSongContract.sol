// SPDX-License-Identifier: MIT
pragma solidity ^0.5.5;
pragma experimental ABIEncoderV2;

contract PublishSongContract {

    struct Song {
        string cid;                // CID (Content identifier) of the song return by IPFS node.

        string name;                // name of the song.
        string[] artists;           // artists who created the song.
        string description;         // description
        uint256 publishedTimestamp; 

        address publisher;          // address of the account publishing the song.

        bool isValid;               // check if field is valid. This field is used to avoid adding 
                                    // songs with same CID.
    }

    address public owner;

    mapping (string => Song) songs;
    uint16 public totalSongs = 0;
    string[] public cids;

    constructor() public {
        owner = msg.sender;
    }

    function addSong(string memory _cid, string memory _name, string[] memory _artists, string memory _description) public {
        require(!songs[_cid].isValid, "Song with given CID already added.");

        totalSongs += 1;
        cids.push(_cid);
        songs[_cid] = Song(_cid, _name, _artists, _description, block.timestamp, msg.sender, true);
    }

    function getSong(string memory _cid) public view returns (Song memory) {
        require(songs[_cid].isValid, "No song with given CID.");
        return songs[_cid];
    }

    function getAllCid() public view returns(string[] memory) {
        return cids;
    }

}