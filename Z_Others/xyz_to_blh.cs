using System;

public class CartesianCoordinate
{
    public double X { get; set; }
    public double Y { get; set; }
    public double Z { get; set; }
}

public class GeodeticCoordinate
{
    public double Latitude { get; set; }
    public double Longitude { get; set; }
    public double Altitude { get; set; }
}

public class CoordinateConverter
{
    private const double WGS84_a = 6378137.0; // 椭球体长半轴
    private const double WGS84_f = 1 / 298.257223563; // 扁率

    public GeodeticCoordinate ConvertCartesianToGeodetic(CartesianCoordinate cartesianCoordinate)
    {
        double X = cartesianCoordinate.X;
        double Y = cartesianCoordinate.Y;
        double Z = cartesianCoordinate.Z;

        double longitude = Math.Atan2(Y, X);

        double R = Math.Sqrt(X * X + Y * Y);
        double latitude = Math.Atan(Z / (R * (1 - GetEccentricitySquared(latitude))));

        double N;
        double delta;
        double altitude;

        do
        {
            N = WGS84_a / Math.Sqrt(1 - GetEccentricitySquared(latitude) * Math.Sin(latitude) * Math.Sin(latitude));
            altitude = R / Math.Cos(latitude) - N;
            delta = altitude / N;
            latitude = Math.Atan(Z / (R * (1 - GetEccentricitySquared(latitude) * N / (N + altitude))));
        } while (Math.Abs(delta) > 1e-5);

        return new GeodeticCoordinate
        {
            Latitude = RadiansToDegrees(latitude),
            Longitude = RadiansToDegrees(longitude),
            Altitude = altitude
        };
    }

    private double GetEccentricitySquared(double latitude)
    {
        double b = WGS84_a * (1 - WGS84_f);
        return (WGS84_a * WGS84_a - b * b) / (WGS84_a * WGS84_a);
    }

    private double RadiansToDegrees(double radians)
    {
        return radians * 180 / Math.PI;
    }
}


public void main(){
    CoordinateConverter converter = new CoordinateConverter();

    CartesianCoordinate cartesianCoordinate = new CartesianCoordinate
    {
        X = 1000.0, // X坐标
        Y = 2000.0, // Y坐标
        Z = 3000.0  // Z坐标
    };

    GeodeticCoordinate geodeticCoordinate = converter.ConvertCartesianToGeodetic(cartesianCoordinate);

    Console.WriteLine($"Latitude: {geodeticCoordinate.Latitude}");
    Console.WriteLine($"Longitude: {geodeticCoordinate.Longitude}");
    Console.WriteLine($"Altitude: {geodeticCoordinate.Altitude}");

}