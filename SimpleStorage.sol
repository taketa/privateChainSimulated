pragma solidity ^0.4.17;

contract SimpleStorage{
    
    mapping (address => User) AllUsers;
    struct MainInfo{
            string visa;
            string nationality;
            uint age;
            string speaks;
    }
    struct MedInfo{
            string[] vitalMedicalConditions;
            string[] medications;
    }
    struct User{
        string      fName;
        string      sName;
        MainInfo    mainInfo;
        MedInfo     medInfo;
    }
    // setters
    function setVisa(string _value){
        AllUsers[msg.sender].mainInfo.visa=_value;
    } 
    function setNationality(string _value){
        AllUsers[msg.sender].mainInfo.nationality=_value;
    } 
    function setAge(uint _value){
        AllUsers[msg.sender].mainInfo.age=_value;
    } 
    function setSpeaks(string _value){
        AllUsers[msg.sender].mainInfo.speaks=_value;
    } 
    function setMedCondition(string _value){
        AllUsers[msg.sender].medInfo.vitalMedicalConditions.push(_value);
    }
    function setMedMedication(string _value){
        AllUsers[msg.sender].medInfo.medications.push(_value);
    }
    
    // getters
    function getVisa() constant returns (string){
        return AllUsers[msg.sender].mainInfo.visa;
    }
    function getNationality() constant returns (string){
        return AllUsers[msg.sender].mainInfo.nationality;
    }
    function getAge() constant returns (uint){
        return AllUsers[msg.sender].mainInfo.age;
    }
    function getSpeaks() constant returns (string){
        AllUsers[msg.sender].mainInfo.speaks;
    } 
    function getMedCondition() returns (string[]){
        AllUsers[msg.sender].medInfo.vitalMedicalConditions;
    }
     function getMedMedication() returns (string[]){
        AllUsers[msg.sender].medInfo.medications;
    }

}

