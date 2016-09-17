import java.util.ArrayList;
import java.util.Collection;

public class LinearRegression {
//	y = a0 + a1x + e
//	e = y − a0 − a1x
	
	private static double a1(double[] x, double[] y) {
		final int n = y.length;
		double sumX = 0,
				sumY = 0,
				sumXY = 0,
				sumXSq = 0;
		for (int i = 0; i < n; i++) {
			sumX += x[i];
			sumY += y[i];
			sumXY += x[i] * y[i];
			sumXSq += Math.pow(x[i], 2);
		}
		return ((n * sumXY) - (sumX * sumY))/((n * sumXSq) - Math.pow(sumX, 2));
	}
	
	private static double a0(double[] x, double[] y, double a1) {
		final int n = y.length;
		double sumx = 0, sumy = 0;
		double _x, _y;
		
		for (int i = 0; i < n; i++) {
			sumx += x[i];
			sumy += y[i];
		}
		
		_x = sumx / n;
		_y = sumy / n;
		
		return _y - a1 * _x;
	}
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		double[] x = {  1,   2,   3,   4,   5,   6,   7};
		double[] y = {0.5, 2.5, 2.0, 4.0, 3.5, 6.0, 5.5};
		double a1 = a1(x, y);
		double a0 = a0(x, y, a1);
		System.out.println(a1);
		System.out.println(a0);
		System.out.println("Line Equation: y = " + a1 + "x + " + a0);
		
		
	}

}
