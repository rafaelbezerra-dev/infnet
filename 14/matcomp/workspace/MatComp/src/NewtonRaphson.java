
public class NewtonRaphson {


	/* 	Achar a raiz de um número n
	 *	- Sendo n = 612
	 *	- É o mesmo que: x² = 612
	 *	- A função que deve ser usada no método
	 *	  de Newton-Raphson é:
	 *	  f(x) = x² - 612
	 *	- A primeira derivada é: f'(x) = 2x
	 *	 		 
	 */



	/* 	Find the square root of a number n
	 *	- Let n = 612
	 *	- Can be represented as: x² = 612
	 *	- The function to use in Newton's method
	 *	  is then:
	 *	  f(x) = x² - 612
	 *	- And the first derivative: f'(x) = 2x
	 *	 		 
	 */
	
	static double f(double x, double c) {
		return Math.pow(x, 2) - c;
	}
	
	static double fprime(double x) {
		return 2 * x;
	}
	
	public static void main(String[] args) {
        // read in the command-line argument				
        final double c = 612.00;//Double.parseDouble(args[0]);
        final double epsilon = 1e-15;		// error tolerance
        final int maxIterations = 20;
//      double x0 = x; 				// initial guess
        double xi = 1; 				// initial guess
        int i;
        
        for (i = 0; i < maxIterations; i++) {
			double x1 = xi - f(xi, c)/fprime(xi);
			
			if(Math.abs(x1 - xi) <= epsilon * Math.abs(x1)) {				
				xi = x1;	
				break;
			}
			
			xi = x1;		
		}
        
        System.out.println(xi);
        
	}

}
