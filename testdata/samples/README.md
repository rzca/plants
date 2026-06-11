# Sample Data Files

These files are extracted/derived from the official USDA Plant Hardiness Zone Map data for testing and development.

## Files

- `grid.hdr` - Header file describing the BIL raster grid format
- `grid.prj` - Projection file (NAD83 geographic coordinates)
- `zones.prj` - Projection file for shapefile data
- `zone_attributes.csv` - Zone attribute lookup table (gridcode -> zone info)
- `terms_of_use.txt` - Official USDA/OSU terms of use

## Grid Data Format

The full grid is a 7025x3105 raster covering the continental US:
- Upper-left corner: -125.017°, 49.933° (roughly northwest WA)
- Cell size: 0.00833° (~1km at mid-latitudes)
- Values: Float32, where value maps to zone via gridcode
- NoData: -9999

## Northern Virginia Test Area

The test zipcodes cover the NoVA region:
- Zone 7a (0 to 5°F): Western Loudoun, Prince William (20101, 20105, 20109, etc.)
- Zone 7b (5 to 10°F): Fairfax, Falls Church, McLean (22003, 22030, 22101, etc.)

Approximate coordinates for testing:
- Reston: 38.9687° N, 77.3411° W (Zone 7a)
- Falls Church: 38.8820° N, 77.1711° W (Zone 7b)
- McLean: 38.9339° N, 77.1773° W (Zone 7b)
