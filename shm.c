#include <stdio.h>
#include <sys/shm.h>
#include <sys/stat.h>
#include <unistd.h>
#include <stdlib.h>
int main()
{
        pid_t PID = fork();
        switch(PID){
                case -1: perror("fork()"); exit(-1);
                case 0:
                        printf("I'm Num1\n");
                        printf("Num1's PID = %d\n",getpid()); break;
                default :
                        printf("I'm Num2\n");
                        printf("Num2's PID = %d\n",getpid());
        }
        int segment_id;
        char* shared_memory;
        const int segment_size = 2048;
        segment_id = shmget(IPC_PRIVATE, segment_size,S_IRUSR | S_IWUSR);
        shared_memory =(char *)shmat(segment_id,NULL,0);
        sprintf(shared_memory,"%d enter in shared memory",segment_id);
        printf("> %s\n",shared_memory);
        if(shmdt(shared_memory) == -1)
                {fprintf(stderr,"Unable to detach\n");}
        shmctl(segment_id,IPC_RMID,NULL);
        return 0;
}