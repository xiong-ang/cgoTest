#include <iostream>
#include <vector>
#include "cppLib.h"

#ifdef __cplusplus
extern "C"
{
#endif

int test()
{
    std::vector<int> v1{1, 2};
    std::vector<int> v2{2, 3};
    v1 = v2;
    v2.push_back(4);
    std::cout << v1.size() << std::endl;
    std::cout << v2.size() << std::endl;
    return 0;
}

#ifdef __cplusplus
}
#endif