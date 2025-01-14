import java.util.Arrays;

public class Jacobi {
	public static float E = 0.000000001f;//Precisão	
	
    public static boolean isDagonallyDominant(float[][] A)
    {
    	float sum = 0f;
    	float diag = 0f;
		for (int i = 0; i < A.length; i++) {
			diag = A[i][i];
			sum = 0;
			for (int j = 0; j < A.length; j++) {
				if(i != j){
					sum += Math.abs(A[i][j]);
				}
			}
			if (diag < sum) {
				return false;
			}
		}
    	return true;
    }
	
	public static boolean shouldStop(float[] x, float[] x0) { 
		for (int i = 0; i < x.length; i++) {
			float sub = x[i] - x0[i];
			if (E < Math.abs(sub)) {
				return false;
			} 
		}
		return true;
	}

	public static void jacobi(float[][] A, float[] b) {
		int i, j, k;
		boolean _return_ = true;
		float sum = 0;
		int max = 1000;
		float[] x = new float[A.length];
		float[] x0 = new float[A.length];
		
		if (isDagonallyDominant(A)){
			for (k = 0; k < max; k++) {
				for (i = 0; i < A.length; i++) {
					sum = 0;
					for (j = 0; j < A.length; j++) {
						if(i != j){
							sum += - (A[i][j] * x0[j]);
						}
					}
					x[i] = (b[i] + sum)/A[i][i];				
				}

				
				System.out.print(k + ") x = ");
				for (i = 0; i < x.length; i++) {
					System.out.print(x[i] + ",  "); 
				}
				
				if(shouldStop(x, x0)) {
					break;
				}
				
				System.out.println("");
				
				// coping
				for (i = 0; i < x0.length; i++) {
					x0[i] = x[i]; 
				}
			}
			
	//		System.out.print("x = ");
	//		for (i = 0; i < x.length; i++) {
	//			System.out.print(x[i] + ",  "); 
	//		}
			
			if (k >= max) {
				_return_ = false;
			}
			else {
				_return_ = true;
			}
		}
		else{
            System.out.println("The system isn't diagonally dominant: " + 
                    "The method cannot guarantee convergence.");
        }
	}
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
//		float[][] A = { 
//				{  4, -1, -1 }, 
//				{ -2,  6,  1 }, 
//				{ -1,  1,  7 } 
//		};
//		float[] b = { 3, 9, -6 };
		
//		float[][] A = { 
//				{ 70,	1,	0 },
//				{ 60,	-1,	1 },
//				{ 40,	0,	-1 }
//		};
//		float[] b = { 636, 518, 307 };
		
		float[][] A = { 
				{ 2,	1},
				{ 5,	7 },
		};
		float[] b = { 11, 13 };
		jacobi(A, b);
	}

}
