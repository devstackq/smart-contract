Pet project: smart-contract 2023
SC: have 3 side - Eth api; Solidity code, Golang - 
-Solidity compiler - deploy Eth-machine;
    Then - (abi -> convert -> go)
-Go-for communicate Eth api

commands:
    solc --bin --abi solidity/todo.sol -o build
    - Solidity compiler build from solidity file -> abi/bin files for create Abigen -> create Go file

    abigen --bin=build/Todo.bin --abi=build/Todo.abi --pkg=todo --out=gen/todo.go
    - abigen from bin/abigen -> generate go file

