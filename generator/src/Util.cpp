#include "Util.hpp"

void Util::Clear(){
    #if defined _WIN32
        system("cls");
    #elif defined (__LINUX__) || defined(__gnu_linux__) || defined(__linux__)
        system("clear");
    #elif defined (__APPLE__)
        system("clear");
    #endif
}

void Util::PrintArray(int *arr, int size){
    cout << "[";
    for (int i = 0; i < size; i++)
    {
        cout << arr[i];
        if (i != size -1)
            cout << ", ";
    }
    cout << "]"    ;
}

string Util::GetNowTime(){
    time_t curr_time;
	tm * curr_tm;
	char date_string[100];
	char time_string[100];
	
	time(&curr_time);
	curr_tm = localtime(&curr_time);
	
	strftime(date_string, 50, "%Y_%d_%m-%H_%M_%S", curr_tm);

    return date_string;
}
