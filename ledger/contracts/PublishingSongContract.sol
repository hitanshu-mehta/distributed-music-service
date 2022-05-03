// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

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

    mapping (string => Song) songs;
    uint16 public totalSongs = 0;
    string[] public cids;

    function addSong(string calldata _cid, string calldata _name, string[] calldata _artists, string calldata _description) public {
        require(!songs[_cid].isValid, "Song with given CID already added.");

        totalSongs += 1;
        cids.push(_cid);
        songs[_cid] = Song(_cid, _name, _artists, _description, block.timestamp, msg.sender, true);
    }

    function getSong(string calldata _cid) public view returns (Song memory) {
        require(songs[_cid].isValid, "No song with given CID.");
        return songs[_cid];
    }

    function getAllCid() public view returns(string[] memory) {
        return cids;
    }

}