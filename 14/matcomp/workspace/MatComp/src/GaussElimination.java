/******************************************************************************
 * Compilation: javac GaussianElimination.java Execution: java
 * GaussianElimination
 * 
 * Gaussian elimination with partial pivoting.
 *
 * % java GaussianElimination -1.0 2.0 2.0
 *
 ******************************************************************************/

public class GaussElimination {
	private static final double EPSILON = 1e-10;

	// Gaussian elimination with partial pivoting
	public static double[] lsolve(double[][] A, double[] b) {
		int N = b.length;

		for (int p = 0; p < N; p++) {

			// find pivot row and swap
			int max = p;
			for (int i = p + 1; i < N; i++) {
				if (Math.abs(A[i][p]) > Math.abs(A[max][p])) {
					max = i;
				}
			}
			double[] temp = A[p];
			A[p] = A[max];
			A[max] = temp;
			double t = b[p];
			b[p] = b[max];
			b[max] = t;

			// singular or nearly singular
			if (Math.abs(A[p][p]) <= EPSILON) {
				throw new RuntimeException("Matrix is singular or nearly singular");
			}

			// pivot within A and b
			for (int i = p + 1; i < N; i++) {
				double alpha = A[i][p] / A[p][p];
				b[i] -= alpha * b[p];
				for (int j = p; j < N; j++) {
					A[i][j] -= alpha * A[p][j];
				}
			}
		}

		// back substitution
		double[] x = new double[N];
		for (int i = N - 1; i >= 0; i--) {
			double sum = 0.0;
			for (int j = i + 1; j < N; j++) {
				sum += A[i][j] * x[j];
			}
			x[i] = (b[i] - sum) / A[i][i];
		}
		return x;
	}

	// sample client
	public static void main(String[] args) {
//		System.out.println(EPSILON);
		int N = 3;
		double[][] A = { { 0, 1, 1 }, { 2, 4, -2 }, { 0, 3, 15 } };
		double[] b = { 4, 2, 36 };
		double[] x = lsolve(A, b);

		// print results
		for (int i = 0; i < N; i++) {
			System.out.println(x[i]);
		}

	}

	// public static void main(String[] args) {
	// // TODO Auto-generated method stub
	//// gauss();
	//
	//
	// double[][] A = {
	// {70, 1, 0},
	// {60, -1, 1},
	// {40, 0, -1},
	// };
	// double[] b = {
	// 636,
	// 518,
	// 307
	// };
	// double[] x = lsolve(A, b);
	//
	// for (double f : x) {
	// System.out.println(f);
	// }
	// }

}
