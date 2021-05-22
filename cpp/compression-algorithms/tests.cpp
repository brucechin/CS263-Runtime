
// ================================================================================================
// -*- C++ -*-
// File: tests.cpp
// Author: Guilherme R. Lampert
// Created on: 15/02/16
// Brief: Basic unit tests for the data compression algorithms.
//
// This source code is in the public domain.
// You are free to do whatever you want with it.
//
// Compile with:
// c++ -std=c++11 -O3 -Wall -Wextra -Weffc++ -Wshadow -pedantic -I.. tests.cpp -o tests
// ================================================================================================



#define HUFFMAN_IMPLEMENTATION
#include "huffman.hpp"


#include <cstdint>
#include <cstring>
#include <iostream>
#include <vector>
#include <chrono>
#include <string>
#include <fstream>
#include <streambuf>




static void Test_Huffman_EncodeDecode(const std::uint8_t * sampleData, const int sampleSize)
{
    int compressedSizeBytes = 0;
    int compressedSizeBits  = 0;
    int repeatTimes = 10;
    std::uint8_t * compressedData = nullptr;
    std::vector<std::uint8_t> uncompressedBuffer(sampleSize, 0);

    // Compress:
    auto startTime = std::chrono::system_clock::now();                                           
    for(int i = 0; i < repeatTimes; ++i){
    huffman::easyEncode(sampleData, sampleSize, &compressedData,
                        &compressedSizeBytes, &compressedSizeBits);
    }

    auto endTime = std::chrono::system_clock::now();                                             
    std::chrono::duration<double> elapsedSeconds = endTime - startTime;                                
    std::cout << ">>> encode completed in " << elapsedSeconds.count()/repeatTimes << " seconds.\n"; 
    std::cout << std::endl;    


    // Restore:
    startTime = std::chrono::system_clock::now();    
    for(int i = 0; i < repeatTimes; ++i){

    const int uncompressedSize = huffman::easyDecode(compressedData, compressedSizeBytes, compressedSizeBits,
                                                    uncompressedBuffer.data(), uncompressedBuffer.size());
    }
    endTime = std::chrono::system_clock::now();                                             
    elapsedSeconds = endTime - startTime;                                
    std::cout << ">>> encode completed in " << elapsedSeconds.count()/repeatTimes << " seconds.\n"; 
    std::cout << std::endl;    
    
    


    // easyEncode() uses HUFFMAN_MALLOC (std::malloc).
    HUFFMAN_MFREE(compressedData);
}

static void Test_Huffman()
{
    std::ifstream t("alice29.txt");
    std::string inputfile((std::istreambuf_iterator<char>(t)),
                    std::istreambuf_iterator<char>());
    std::vector<uint8_t> myVector(inputfile.begin(), inputfile.end());
    uint8_t *p = &myVector[0];
    Test_Huffman_EncodeDecode(p, inputfile.size());

}



int main()
{
    std::cout << "\nRunning unit tests for the compression algorithms...\n\n";

    Test_Huffman();

}

// ========================================================

