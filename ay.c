#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>

int main(int argc, char** argv) {
    // Initialize MPI
    MPI_Init(NULL, NULL);

    // Get the number of processes
    int num_procs;
    MPI_Comm_size(MPI_COMM_WORLD, &num_procs);

    // Get the rank of the process
    int rank;
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    // Initialize variables for the number of ripe and unripe fruits
    int ripe_count = 0, unripe_count = 0;

    // The farmer has 3 acres of land in three different villages
    for (int i = 0; i < 3; i++) {
        // The robots are deployed for counting by assigning one row to each robot
        for (int j = 0; j < 4; j++) {
            // The robot with rank `j` counts the fruits in its row
            if (rank == j) {
                // The robot inspects the fruits and determines whether they are ripe or not
                for (int k = 0; k < 16; k++) {
                    int r = rand();
                    int fruit_status = r > RAND_MAX/2 ? 1 : 0;// determine whether the fruit is ripe or not
                        if (fruit_status == 1) {
                            ripe_count++;
                        }
                        else {
                            unripe_count++;
                        }
                }
            }
            // The robots send a message to the managerial process containing the running count of the number of fruits that are ripe & number that are not yet ripe for every set of four trees
            int ripe_count_global, unripe_count_global;
            MPI_Reduce(&ripe_count, &ripe_count_global, 1, MPI_INT, MPI_SUM, 0, MPI_COMM_WORLD);
            MPI_Reduce(&unripe_count, &unripe_count_global, 1, MPI_INT, MPI_SUM, 0, MPI_COMM_WORLD);

            // The managerial process prints the global counts of ripe and unripe fruits
            if (rank == 0) {
                printf("Total number of ripe fruits: %d\n", ripe_count_global);
                printf("Total number of unripe fruits: %d\n", unripe_count_global);
            }

            // The robots are moved together to the next piece of land and this process of counting is repeated
            ripe_count = 0;
            unripe_count = 0;
        }
    }

    // Finalize MPI
    MPI_Finalize();
}
