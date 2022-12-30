#include <mpi.h>
#include <stdio.h>
#include <stdlib.h>

const int NUM_ACRES = 3;
const int NUM_ROWS_PER_ACRE = 4;
const int NUM_TREES_PER_ROW = 16;

int main(int argc, char** argv) {
  MPI_Init(NULL, NULL);

  // Get the number of processes
  int size;
  MPI_Comm_size(MPI_COMM_WORLD, &size);

  // Get the rank of the process
  int rank;
  MPI_Comm_rank(MPI_COMM_WORLD, &rank);

  // Initialize variables for the number of ripe and unripe fruits in each acre
  int ripe_fruits[NUM_ACRES] = {0};
  int unripe_fruits[NUM_ACRES] = {0};

  // Divide the work of inspecting the fruits among the processes
  for (int i = 0; i < NUM_ACRES; i++) {
    // Calculate the number of rows assigned to this process
    int rows_assigned = NUM_ROWS_PER_ACRE / size;
    // Calculate the starting and ending row indices for this process
    int start_row = rank * rows_assigned;
    int end_row = start_row + rows_assigned;

    // Inspect the fruits in the assigned rows
    for (int j = start_row; j < end_row; j++) {
      // Simulate inspecting the fruits in this row
      int num_ripe = rand() % NUM_TREES_PER_ROW;
      int num_unripe = NUM_TREES_PER_ROW - num_ripe;
      // Update the count for this acre
      ripe_fruits[i] += num_ripe;
      unripe_fruits[i] += num_unripe;
    }

    // Send the count for this acre to the managerial process
    if (rank != 0) {
      MPI_Send(&ripe_fruits[i], 1, MPI_INT, 0, i, MPI_COMM_WORLD);
      MPI_Send(&unripe_fruits[i], 1, MPI_INT, 0, i, MPI_COMM_WORLD);
    }
  }

  // Managerial process receives the count from each process and updates the total count for each acre
  if (rank == 0) {
    int total_ripe_fruits[NUM_ACRES] = {0};
    int total_unripe_fruits[NUM_ACRES] = {0};
    for (int i = 1; i < size; i++) {
      for (int j = 0; j < NUM_ACRES; j++) {
        int ripe, unripe;
        MPI_Recv(&ripe, 1, MPI_INT, i, j, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
	MPI_Recv(&unripe, 1, MPI_INT, i, j, MPI_COMM_WORLD, MPI_STATUS_IGNORE);
        total_ripe_fruits[j] += ripe;
        total_unripe_fruits[j] += unripe;
      }
    }

    // Print the final count for each acre
    for (int i = 0; i < NUM_ACRES; i++) {
      printf("Acre %d: %d ripe fruits, %d unripe fruits\n", i, total_ripe_fruits[i], total_unripe_fruits[i]);
    }
  }

  // Finalize the MPI environment.
  MPI_Finalize();
}

