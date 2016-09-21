
public class BetterNewtonRaphson {


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
//		return Math.pow(x, 2) - c;
		return Math.pow(x, 10) - c;
	}
	
	static double fprime(double x) {
//		return 2 * x;
		return 10 * Math.pow(x, 9);
	}
	
	public static void main(String[] args) {
        // read in the command-line argument				
//        final double c = Double.parseDouble(args[0]);
        final double c = 1;
        final double epsilon = 1e-15;		// error tolerance
        final int maxIterations = 100;
        final boolean debug = args.length > 2 ? Boolean.parseBoolean(args[2]) : false; 
//        double xi = args.length > 1 ? Double.parseDouble(args[1]) : 1; 				// initial guess
        double xi = 0.5;
        
        for (int i = 0; i < maxIterations; i++) {
			double x = xi - f(xi, c)/fprime(xi);
			
			if (debug) {
				System.out.println(i + " - xi: " + xi + "; x: " + x + "; err: " + (epsilon * Math.abs(x)));
			}
			
			if(Math.abs(x - xi) <= epsilon * Math.abs(x)) {				
				xi = x;	
				break;
			}
			
			xi = x;		
		}
        
        System.out.println("Root is: " + xi);
        
	}

}
