// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 < 0.9.0;

contract Todo{

    Task[] tasks;
    
    address public owner;
    
    struct Task{
        // string id;
        string content;
        bool statusl;
        string author;
    }
    modifier isOwner(){//like decorator
            require(owner == msg.sender);
            _;
    }
        constructor(){
            owner = msg.sender;
        }
        function add(string memory _content, string  memory _author)public isOwner{
            tasks.push(Task(_content, false, _author));
        }
        
        function get(uint _id )public  view returns (Task memory){
            return tasks[_id];
        }

        function list() public  view returns (Task[] memory){
            return tasks;
        }

        function update(uint _id, string memory _content)public {
            tasks[_id].content= _content;
        }

        function remove(uint _id)public {

            delete   tasks[_id];
        }

}