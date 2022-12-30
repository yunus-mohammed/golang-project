#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>

#define NUM_HOURS 3 // number of hours each restaurant works

// Data structure for restaurant data
typedef struct {
  int num_customers;
  int num_idlis;
  int num_vadas;
  int num_dosas;
  int num_rava_idlis;
} restaurant_data;

int main(int argc, char** argv) {
  MPI_Init(NULL, NULL);

  int world_size;
  MPI_Comm_size(MPI_COMM_WORLD, &world_size);

  int world_rank;
  MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);

  // Master process (rank 0) collects data from all restaurants
  if (world_rank == 0) {
    restaurant_data total = {0};

    // Receive data from each restaurant
    for (int i = 1; i < world_size; i++) {
      restaurant_data data;
      MPI_Recv(&data, sizeof(restaurant_data), MPI_BYTE, i, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
      total.num_customers += data.num_customers;
      total.num_idlis += data.num_idlis;
      total.num_vadas += data.num_vadas;
      total.num_dosas += data.num_dosas;
      total.num_rava_idlis += data.num_rava_idlis;
    }

    // Send the total data to all restaurants
    for (int i = 1; i < world_size; i++) {
      MPI_Send(&total, sizeof(restaurant_data), MPI_BYTE, i, 0, MPI_COMM_WORLD);
    }
  } else { // Other processes (restaurants) send data to the master process
    for (int i = 0; i < NUM_HOURS; i++) {
      // Simulate collecting data by generating random numbers
      restaurant_data data = {
        rand(), rand(), rand(), rand(), rand()
      };
      MPI_Send(&data, sizeof(restaurant_data), MPI_BYTE, 0, 0, MPI_COMM_WORLD);
    }

    // Receive the total data from the master process
    restaurant_data total;
    MPI_Recv(&total, sizeof(restaurant_data), MPI_BYTE, 0, 0, MPI_COMM_WORLD, MPI_STATUS_IGNORE);

    // Print the total data
    printf("Restaurant %d received total data:\n", world_rank);
    printf("Number of customers: %d\n", total.num_customers);
    printf("Number of rice idlis: %d\n", total.num_idlis);
    printf("Number of vadas: %d\n", total.num_vadas);
    printf("Number of dosas: %d\n", total.num_dosas);
  }
    MPI_Finalize();
}
