package uom_checker

const uomsXML = `
<?xml version="1.0" encoding="UTF-8" ?>
<Data>
    <UOM>
    <code>VL</code>
    <desc>VIAL</desc>
  </UOM>
  <UOM>
    <code>AA</code>
    <desc>Ball</desc>
  </UOM>
  <UOM>
    <code>AB</code>
    <desc>Bulk Pack</desc>
  </UOM>
  <UOM>
    <code>AC</code>
    <desc>Acre</desc>
  </UOM>
  <UOM>
    <code>AD</code>
    <desc>Bytes</desc>
  </UOM>
  <UOM>
    <code>AE</code>
    <desc>Amperes per Meter</desc>
  </UOM>
  <UOM>
    <code>AF</code>
    <desc>Centigram</desc>
  </UOM>
  <UOM>
    <code>AG</code>
    <desc>Angstrom</desc>
  </UOM>
  <UOM>
    <code>AH</code>
    <desc>Additional Minutes</desc>
  </UOM>
  <UOM>
    <code>AI</code>
    <desc>Average Minutes Per Call</desc>
  </UOM>
  <UOM>
    <code>AJ</code>
    <desc>Cop</desc>
  </UOM>
  <UOM>
    <code>AK</code>
    <desc>Fathom</desc>
  </UOM>
  <UOM>
    <code>AL</code>
    <desc>Access Lines</desc>
  </UOM>
  <UOM>
    <code>AM</code>
    <desc>Ampoule</desc>
  </UOM>
  <UOM>
    <code>AN</code>
    <desc>Minutes or Messages</desc>
  </UOM>
  <UOM>
    <code>AO</code>
    <desc>Ampere-turn</desc>
  </UOM>
  <UOM>
    <code>AP</code>
    <desc>Aluminum Pounds Only</desc>
  </UOM>
  <UOM>
    <code>AQ</code>
    <desc>Anti-hemophilic Factor (AHF) Units</desc>
  </UOM>
  <UOM>
    <code>AR</code>
    <desc>Suppository</desc>
  </UOM>
  <UOM>
    <code>AS</code>
    <desc>Assortment</desc>
  </UOM>
  <UOM>
    <code>AT</code>
    <desc>Atmosphere</desc>
  </UOM>
  <UOM>
    <code>AU</code>
    <desc>Ocular Insert System</desc>
  </UOM>
  <UOM>
    <code>AV</code>
    <desc>Capsule</desc>
  </UOM>
  <UOM>
    <code>AW</code>
    <desc>Powder-Filled Vials</desc>
  </UOM>
  <UOM>
    <code>AX</code>
    <desc>Twenty</desc>
  </UOM>
  <UOM>
    <code>AY</code>
    <desc>Assembly</desc>
  </UOM>
  <UOM>
    <code>AZ</code>
    <desc>British Thermal Units (BTUs) per Pound</desc>
  </UOM>
  <UOM>
    <code>A8</code>
    <desc>Dollars per Hours</desc>
  </UOM>
  <UOM>
    <code>BA</code>
    <desc>Bale</desc>
  </UOM>
  <UOM>
    <code>BB</code>
    <desc>Base Box</desc>
  </UOM>
  <UOM>
    <code>BC</code>
    <desc>Bucket</desc>
  </UOM>
  <UOM>
    <code>BD</code>
    <desc>Bundle</desc>
  </UOM>
  <UOM>
    <code>BE</code>
    <desc>Beam</desc>
  </UOM>
  <UOM>
    <code>BF</code>
    <desc>Board Feet</desc>
  </UOM>
  <UOM>
    <code>BG</code>
    <desc>Bag</desc>
  </UOM>
  <UOM>
    <code>BH</code>
    <desc>Brush</desc>
  </UOM>
  <UOM>
    <code>BI</code>
    <desc>Bar</desc>
  </UOM>
  <UOM>
    <code>BJ</code>
    <desc>Band</desc>
  </UOM>
  <UOM>
    <code>BK</code>
    <desc>Book</desc>
  </UOM>
  <UOM>
    <code>BL</code>
    <desc>Block</desc>
  </UOM>
  <UOM>
    <code>BM</code>
    <desc>Bolt</desc>
  </UOM>
  <UOM>
    <code>BN</code>
    <desc>Bulk</desc>
  </UOM>
  <UOM>
    <code>BO</code>
    <desc>Bottle</desc>
  </UOM>
  <UOM>
    <code>BP</code>
    <desc>100 Board Feet</desc>
  </UOM>
  <UOM>
    <code>BQ</code>
    <desc>Brake horse power</desc>
  </UOM>
  <UOM>
    <code>BR</code>
    <desc>Barrel</desc>
  </UOM>
  <UOM>
    <code>BS</code>
    <desc>Basket</desc>
  </UOM>
  <UOM>
    <code>BT</code>
    <desc>Belt</desc>
  </UOM>
  <UOM>
    <code>BU</code>
    <desc>Bushel</desc>
  </UOM>
  <UOM>
    <code>BV</code>
    <desc>Bushel, Dry Imperial</desc>
  </UOM>
  <UOM>
    <code>BW</code>
    <desc>Base Weight</desc>
  </UOM>
  <UOM>
    <code>BX</code>
    <desc>Box</desc>
  </UOM>
  <UOM>
    <code>BY</code>
    <desc>British Thermal Unit (BTU)</desc>
  </UOM>
  <UOM>
    <code>BZ</code>
    <desc>Million BTU's</desc>
  </UOM>
  <UOM>
    <code>B0</code>
    <desc>British Thermal Units (BTUs) per Cubic Foot</desc>
  </UOM>
  <UOM>
    <code>B2</code>
    <desc>Bunks</desc>
  </UOM>
  <UOM>
    <code>B3</code>
    <desc>Batting Pound</desc>
  </UOM>
  <UOM>
    <code>B4</code>
    <desc>Barrel, Imperial</desc>
  </UOM>
  <UOM>
    <code>B5</code>
    <desc>Billet</desc>
  </UOM>
  <UOM>
    <code>B6</code>
    <desc>Bun</desc>
  </UOM>
  <UOM>
    <code>B7</code>
    <desc>Cycles</desc>
  </UOM>
  <UOM>
    <code>B8</code>
    <desc>Board</desc>
  </UOM>
  <UOM>
    <code>B9</code>
    <desc>Batt</desc>
  </UOM>
  <UOM>
    <code>CA</code>
    <desc>Case</desc>
  </UOM>
  <UOM>
    <code>CB</code>
    <desc>Carboy</desc>
  </UOM>
  <UOM>
    <code>CC</code>
    <desc>Cubic Centimeter</desc>
  </UOM>
  <UOM>
    <code>CD</code>
    <desc>Carat</desc>
  </UOM>
  <UOM>
    <code>CE</code>
    <desc>Centigrade, Celsius</desc>
  </UOM>
  <UOM>
    <code>CF</code>
    <desc>Cubic Feet</desc>
  </UOM>
  <UOM>
    <code>CG</code>
    <desc>Card</desc>
  </UOM>
  <UOM>
    <code>CH</code>
    <desc>Container</desc>
  </UOM>
  <UOM>
    <code>CI</code>
    <desc>Cubic Inches</desc>
  </UOM>
  <UOM>
    <code>CJ</code>
    <desc>Cone</desc>
  </UOM>
  <UOM>
    <code>CK</code>
    <desc>Connector</desc>
  </UOM>
  <UOM>
    <code>CL</code>
    <desc>Cylinder</desc>
  </UOM>
  <UOM>
    <code>CM</code>
    <desc>Centimeter</desc>
  </UOM>
  <UOM>
    <code>CN</code>
    <desc>Can</desc>
  </UOM>
  <UOM>
    <code>CO</code>
    <desc>Cubic Meters (Net)</desc>
  </UOM>
  <UOM>
    <code>CP</code>
    <desc>Crate</desc>
  </UOM>
  <UOM>
    <code>CQ</code>
    <desc>Cartridge</desc>
  </UOM>
  <UOM>
    <code>CR</code>
    <desc>Cubic Meter</desc>
  </UOM>
  <UOM>
    <code>CS</code>
    <desc>Cassette</desc>
  </UOM>
  <UOM>
    <code>CT</code>
    <desc>Carton</desc>
  </UOM>
  <UOM>
    <code>CU</code>
    <desc>Cup</desc>
  </UOM>
  <UOM>
    <code>CV</code>
    <desc>Cover</desc>
  </UOM>
  <UOM>
    <code>CW</code>
    <desc>Hundred Pounds (CWT)</desc>
  </UOM>
  <UOM>
    <code>CX</code>
    <desc>Coil</desc>
  </UOM>
  <UOM>
    <code>CY</code>
    <desc>Cubic Yard</desc>
  </UOM>
  <UOM>
    <code>CZ</code>
    <desc>Combo</desc>
  </UOM>
  <UOM>
    <code>C0</code>
    <desc>Calls</desc>
  </UOM>
  <UOM>
    <code>C1</code>
    <desc>Composite Product Pounds (Total Weight)</desc>
  </UOM>
  <UOM>
    <code>C2</code>
    <desc>Carset</desc>
  </UOM>
  <UOM>
    <code>C3</code>
    <desc>Centiliter</desc>
  </UOM>
  <UOM>
    <code>C4</code>
    <desc>Carload</desc>
  </UOM>
  <UOM>
    <code>C5</code>
    <desc>Cost</desc>
  </UOM>
  <UOM>
    <code>C6</code>
    <desc>Cell</desc>
  </UOM>
  <UOM>
    <code>C7</code>
    <desc>Centipoise (CPS)</desc>
  </UOM>
  <UOM>
    <code>C8</code>
    <desc>Cubic Decimeter</desc>
  </UOM>
  <UOM>
    <code>C9</code>
    <desc>Coil Group</desc>
  </UOM>
  <UOM>
    <code>DA</code>
    <desc>Days</desc>
  </UOM>
  <UOM>
    <code>DB</code>
    <desc>Dry Pounds</desc>
  </UOM>
  <UOM>
    <code>DC</code>
    <desc>Disk (Disc)</desc>
  </UOM>
  <UOM>
    <code>DD</code>
    <desc>Degree</desc>
  </UOM>
  <UOM>
    <code>DE</code>
    <desc>Deal</desc>
  </UOM>
  <UOM>
    <code>DF</code>
    <desc>Dram</desc>
  </UOM>
  <UOM>
    <code>DG</code>
    <desc>Decigram</desc>
  </UOM>
  <UOM>
    <code>DH</code>
    <desc>Miles</desc>
  </UOM>
  <UOM>
    <code>DJ</code>
    <desc>Decagram</desc>
  </UOM>
  <UOM>
    <code>DK</code>
    <desc>Kilometers</desc>
  </UOM>
  <UOM>
    <code>DL</code>
    <desc>Deciliter</desc>
  </UOM>
  <UOM>
    <code>DM</code>
    <desc>Decimeter</desc>
  </UOM>
  <UOM>
    <code>DN</code>
    <desc>Deci Newton-Meter</desc>
  </UOM>
  <UOM>
    <code>DO</code>
    <desc>Dollars, U.S.</desc>
  </UOM>
  <UOM>
    <code>DP</code>
    <desc>Dozen Pair</desc>
  </UOM>
  <UOM>
    <code>DQ</code>
    <desc>Data Records</desc>
  </UOM>
  <UOM>
    <code>DR</code>
    <desc>Drum</desc>
  </UOM>
  <UOM>
    <code>DS</code>
    <desc>Display</desc>
  </UOM>
  <UOM>
    <code>DT</code>
    <desc>Dry Ton</desc>
  </UOM>
  <UOM>
    <code>DU</code>
    <desc>Dyne</desc>
  </UOM>
  <UOM>
    <code>DW</code>
    <desc>Calendar Days</desc>
  </UOM>
  <UOM>
    <code>DX</code>
    <desc>Dynes per Centimeter</desc>
  </UOM>
  <UOM>
    <code>DY</code>
    <desc>Directory Books</desc>
  </UOM>
  <UOM>
    <code>DZ</code>
    <desc>Dozen</desc>
  </UOM>
  <UOM>
    <code>D2</code>
    <desc>Shares</desc>
  </UOM>
  <UOM>
    <code>D3</code>
    <desc>Square Decimeter</desc>
  </UOM>
  <UOM>
    <code>D5</code>
    <desc>Kilogram Per Square Centimeter</desc>
  </UOM>
  <UOM>
    <code>D8</code>
    <desc>Draize Score</desc>
  </UOM>
  <UOM>
    <code>D9</code>
    <desc>Dyne per Square Centimeter</desc>
  </UOM>
  <UOM>
    <code>EA</code>
    <desc>Each</desc>
  </UOM>
  <UOM>
    <code>EB</code>
    <desc>Electronic Mail Boxes</desc>
  </UOM>
  <UOM>
    <code>EC</code>
    <desc>Each per Month</desc>
  </UOM>
  <UOM>
    <code>ED</code>
    <desc>Inches, Decimal--Nominal</desc>
  </UOM>
  <UOM>
    <code>EE</code>
    <desc>Employees</desc>
  </UOM>
  <UOM>
    <code>EF</code>
    <desc>Inches, Fraction-Nominal</desc>
  </UOM>
  <UOM>
    <code>EG</code>
    <desc>Double-time Hours</desc>
  </UOM>
  <UOM>
    <code>EH</code>
    <desc>Knots</desc>
  </UOM>
  <UOM>
    <code>EJ</code>
    <desc>Locations</desc>
  </UOM>
  <UOM>
    <code>EM</code>
    <desc>Inches, Decimal-Minimum</desc>
  </UOM>
  <UOM>
    <code>EP</code>
    <desc>Eleven pack</desc>
  </UOM>
  <UOM>
    <code>EQ</code>
    <desc>Equivalent Gallons</desc>
  </UOM>
  <UOM>
    <code>EV</code>
    <desc>Envelope</desc>
  </UOM>
  <UOM>
    <code>EX</code>
    <desc>Feet, Inches and Fraction</desc>
  </UOM>
  <UOM>
    <code>EY</code>
    <desc>Feet, Inches and Decimal</desc>
  </UOM>
  <UOM>
    <code>EZ</code>
    <desc>Feet and Decimal</desc>
  </UOM>
  <UOM>
    <code>E1</code>
    <desc>Hectometer</desc>
  </UOM>
  <UOM>
    <code>E3</code>
    <desc>Inches, Fraction--Average</desc>
  </UOM>
  <UOM>
    <code>E4</code>
    <desc>Inches, Fraction--Minimum</desc>
  </UOM>
  <UOM>
    <code>E5</code>
    <desc>Inches, Fraction--Actual</desc>
  </UOM>
  <UOM>
    <code>E7</code>
    <desc>Inches, Decimal--Average</desc>
  </UOM>
  <UOM>
    <code>E8</code>
    <desc>Inches, Decimal--Actual</desc>
  </UOM>
  <UOM>
    <code>E9</code>
    <desc>English, (Feet, Inches)</desc>
  </UOM>
  <UOM>
    <code>FA</code>
    <desc>Fahrenheit</desc>
  </UOM>
  <UOM>
    <code>FB</code>
    <desc>Fields</desc>
  </UOM>
  <UOM>
    <code>FC</code>
    <desc>1000 Cubic Feet</desc>
  </UOM>
  <UOM>
    <code>FD</code>
    <desc>Million Particles per Cubic Foot</desc>
  </UOM>
  <UOM>
    <code>FE</code>
    <desc>Track Foot</desc>
  </UOM>
  <UOM>
    <code>FF</code>
    <desc>Hundred Cubic Meters</desc>
  </UOM>
  <UOM>
    <code>FG</code>
    <desc>Transdermal Patch</desc>
  </UOM>
  <UOM>
    <code>FH</code>
    <desc>Micromolar</desc>
  </UOM>
  <UOM>
    <code>FJ</code>
    <desc>Sizing Factor</desc>
  </UOM>
  <UOM>
    <code>FK</code>
    <desc>Fibers</desc>
  </UOM>
  <UOM>
    <code>FM</code>
    <desc>Million Cubic Feet</desc>
  </UOM>
  <UOM>
    <code>FO</code>
    <desc>Fluid Ounce</desc>
  </UOM>
  <UOM>
    <code>FP</code>
    <desc>Pounds per Sq. Ft.</desc>
  </UOM>
  <UOM>
    <code>FR</code>
    <desc>Feet Per Minute</desc>
  </UOM>
  <UOM>
    <code>FS</code>
    <desc>Feet Per Second</desc>
  </UOM>
  <UOM>
    <code>FT</code>
    <desc>Foot</desc>
  </UOM>
  <UOM>
    <code>FZ</code>
    <desc>Fluid Ounce (Imperial)</desc>
  </UOM>
  <UOM>
    <code>F1</code>
    <desc>Thousand Cubic Feet Per Day</desc>
  </UOM>
  <UOM>
    <code>F2</code>
    <desc>International Unit</desc>
  </UOM>
  <UOM>
    <code>F3</code>
    <desc>Equivalent</desc>
  </UOM>
  <UOM>
    <code>F4</code>
    <desc>Minim</desc>
  </UOM>
  <UOM>
    <code>F5</code>
    <desc>MOL</desc>
  </UOM>
  <UOM>
    <code>F6</code>
    <desc>Price Per Share</desc>
  </UOM>
  <UOM>
    <code>F9</code>
    <desc>Fibers per Cubic Centimeter of Air</desc>
  </UOM>
  <UOM>
    <code>GA</code>
    <desc>Gallon</desc>
  </UOM>
  <UOM>
    <code>GB</code>
    <desc>Gallons/Day</desc>
  </UOM>
  <UOM>
    <code>GC</code>
    <desc>Grams per 100 Grams</desc>
  </UOM>
  <UOM>
    <code>GD</code>
    <desc>Gross Barrels</desc>
  </UOM>
  <UOM>
    <code>GE</code>
    <desc>Pounds per Gallon</desc>
  </UOM>
  <UOM>
    <code>GF</code>
    <desc>Grams per 100 Centimeters</desc>
  </UOM>
  <UOM>
    <code>GG</code>
    <desc>Great Gross (Dozen Gross)</desc>
  </UOM>
  <UOM>
    <code>GH</code>
    <desc>Half Gallon</desc>
  </UOM>
  <UOM>
    <code>GI</code>
    <desc>Imperial Gallons</desc>
  </UOM>
  <UOM>
    <code>GJ</code>
    <desc>Grams per Milliliter</desc>
  </UOM>
  <UOM>
    <code>GK</code>
    <desc>Grams per Kilogram</desc>
  </UOM>
  <UOM>
    <code>GL</code>
    <desc>Grams per Liter</desc>
  </UOM>
  <UOM>
    <code>GM</code>
    <desc>Grams per Sq. Meter</desc>
  </UOM>
  <UOM>
    <code>GN</code>
    <desc>Gross Gallons</desc>
  </UOM>
  <UOM>
    <code>GO</code>
    <desc>Milligrams per Square Meter</desc>
  </UOM>
  <UOM>
    <code>GP</code>
    <desc>Milligrams per Cubic Meter</desc>
  </UOM>
  <UOM>
    <code>GQ</code>
    <desc>Micrograms per Cubic Meter</desc>
  </UOM>
  <UOM>
    <code>GR</code>
    <desc>Gram</desc>
  </UOM>
  <UOM>
    <code>GS</code>
    <desc>Gross</desc>
  </UOM>
  <UOM>
    <code>GT</code>
    <desc>Gross Kilogram</desc>
  </UOM>
  <UOM>
    <code>GU</code>
    <desc>Gauss per Oersteds</desc>
  </UOM>
  <UOM>
    <code>GV</code>
    <desc>Gigajoules</desc>
  </UOM>
  <UOM>
    <code>GW</code>
    <desc>Gallons Per Thousand Cubic Feet</desc>
  </UOM>
  <UOM>
    <code>GX</code>
    <desc>Grain</desc>
  </UOM>
  <UOM>
    <code>GY</code>
    <desc>Gross Yard</desc>
  </UOM>
  <UOM>
    <code>GZ</code>
    <desc>Gage Systems</desc>
  </UOM>
  <UOM>
    <code>G2</code>
    <desc>U.S. Gallons Per Minute</desc>
  </UOM>
  <UOM>
    <code>G3</code>
    <desc>Imperial Gallons Per Minute</desc>
  </UOM>
  <UOM>
    <code>G4</code>
    <desc>Gigabecquerel</desc>
  </UOM>
  <UOM>
    <code>G5</code>
    <desc>Gill (Imperial)</desc>
  </UOM>
  <UOM>
    <code>G7</code>
    <desc>Microfiche Sheet</desc>
  </UOM>
  <UOM>
    <code>HA</code>
    <desc>Hank</desc>
  </UOM>
  <UOM>
    <code>HB</code>
    <desc>Hundred Boxes</desc>
  </UOM>
  <UOM>
    <code>HC</code>
    <desc>Hundred Count</desc>
  </UOM>
  <UOM>
    <code>HD</code>
    <desc>Half Dozen</desc>
  </UOM>
  <UOM>
    <code>HE</code>
    <desc>Hundredth of a Carat</desc>
  </UOM>
  <UOM>
    <code>HF</code>
    <desc>Hundred Feet</desc>
  </UOM>
  <UOM>
    <code>HG</code>
    <desc>Hectogram</desc>
  </UOM>
  <UOM>
    <code>HH</code>
    <desc>Hundred Cubic Feet</desc>
  </UOM>
  <UOM>
    <code>HI</code>
    <desc>Hundred Sheets</desc>
  </UOM>
  <UOM>
    <code>HK</code>
    <desc>Hundred Kilograms</desc>
  </UOM>
  <UOM>
    <code>HL</code>
    <desc>Hundred Feet - Linear</desc>
  </UOM>
  <UOM>
    <code>HM</code>
    <desc>Miles Per Hour</desc>
  </UOM>
  <UOM>
    <code>HN</code>
    <desc>Millimeters of Mercury</desc>
  </UOM>
  <UOM>
    <code>HO</code>
    <desc>Hundred Troy Ounces</desc>
  </UOM>
  <UOM>
    <code>HP</code>
    <desc>Millimeter H20</desc>
  </UOM>
  <UOM>
    <code>HQ</code>
    <desc>Hectare</desc>
  </UOM>
  <UOM>
    <code>HR</code>
    <desc>Hours</desc>
  </UOM>
  <UOM>
    <code>HS</code>
    <desc>Hundred Square Feet</desc>
  </UOM>
  <UOM>
    <code>HT</code>
    <desc>Half Hour</desc>
  </UOM>
  <UOM>
    <code>HU</code>
    <desc>Hundred</desc>
  </UOM>
  <UOM>
    <code>HV</code>
    <desc>Hundred Weight (Short)</desc>
  </UOM>
  <UOM>
    <code>HW</code>
    <desc>Hundred Weight (Long)</desc>
  </UOM>
  <UOM>
    <code>HY</code>
    <desc>Hundred Yards</desc>
  </UOM>
  <UOM>
    <code>HZ</code>
    <desc>Hertz</desc>
  </UOM>
  <UOM>
    <code>H1</code>
    <desc>Half Pages - Electronic</desc>
  </UOM>
  <UOM>
    <code>H2</code>
    <desc>Half Liter</desc>
  </UOM>
  <UOM>
    <code>H4</code>
    <desc>Hectoliter</desc>
  </UOM>
  <UOM>
    <code>IA</code>
    <desc>Inch Pound</desc>
  </UOM>
  <UOM>
    <code>IB</code>
    <desc>Inches Per Second (Vibration Velocity)</desc>
  </UOM>
  <UOM>
    <code>IC</code>
    <desc>Counts per Inch</desc>
  </UOM>
  <UOM>
    <code>IE</code>
    <desc>Person</desc>
  </UOM>
  <UOM>
    <code>IF</code>
    <desc>Inches of Water</desc>
  </UOM>
  <UOM>
    <code>IH</code>
    <desc>Inhaler</desc>
  </UOM>
  <UOM>
    <code>II</code>
    <desc>Column-Inches</desc>
  </UOM>
  <UOM>
    <code>IK</code>
    <desc>Peaks per Inch (PPI)</desc>
  </UOM>
  <UOM>
    <code>IL</code>
    <desc>Inches per Minute</desc>
  </UOM>
  <UOM>
    <code>IM</code>
    <desc>Impressions</desc>
  </UOM>
  <UOM>
    <code>IN</code>
    <desc>Inch</desc>
  </UOM>
  <UOM>
    <code>IP</code>
    <desc>Insurance Policy</desc>
  </UOM>
  <UOM>
    <code>IT</code>
    <desc>Counts per Centimeter</desc>
  </UOM>
  <UOM>
    <code>IU</code>
    <desc>Inches Per Second (Linear Speed)</desc>
  </UOM>
  <UOM>
    <code>IV</code>
    <desc>Inches Per Second Per Second (Acceleration)</desc>
  </UOM>
  <UOM>
    <code>IW</code>
    <desc>Inches Per Second Per Second (Vibration Acceleration)</desc>
  </UOM>
  <UOM>
    <code>JA</code>
    <desc>Job</desc>
  </UOM>
  <UOM>
    <code>JB</code>
    <desc>Jumbo</desc>
  </UOM>
  <UOM>
    <code>JE</code>
    <desc>Joule Per Kelvin</desc>
  </UOM>
  <UOM>
    <code>JG</code>
    <desc>Joule per Gram</desc>
  </UOM>
  <UOM>
    <code>JK</code>
    <desc>Mega Joule per Kilogram</desc>
  </UOM>
  <UOM>
    <code>JM</code>
    <desc>Megajoule/Cubic Meter</desc>
  </UOM>
  <UOM>
    <code>JO</code>
    <desc>Joint</desc>
  </UOM>
  <UOM>
    <code>JR</code>
    <desc>Jar</desc>
  </UOM>
  <UOM>
    <code>JU</code>
    <desc>Jug</desc>
  </UOM>
  <UOM>
    <code>J2</code>
    <desc>Joule Per Kilogram</desc>
  </UOM>
  <UOM>
    <code>KA</code>
    <desc>Cake</desc>
  </UOM>
  <UOM>
    <code>KB</code>
    <desc>Kilocharacters</desc>
  </UOM>
  <UOM>
    <code>KC</code>
    <desc>Kilograms per Cubic Meter</desc>
  </UOM>
  <UOM>
    <code>KD</code>
    <desc>Kilograms Decimal</desc>
  </UOM>
  <UOM>
    <code>KE</code>
    <desc>Keg</desc>
  </UOM>
  <UOM>
    <code>KF</code>
    <desc>Kilopackets</desc>
  </UOM>
  <UOM>
    <code>KG</code>
    <desc>Kilogram</desc>
  </UOM>
  <UOM>
    <code>KH</code>
    <desc>Kilowatt Hour</desc>
  </UOM>
  <UOM>
    <code>KI</code>
    <desc>Kilograms/Millimeter Width</desc>
  </UOM>
  <UOM>
    <code>KJ</code>
    <desc>Kilosegments</desc>
  </UOM>
  <UOM>
    <code>KL</code>
    <desc>Kilograms/Meter</desc>
  </UOM>
  <UOM>
    <code>KM</code>
    <desc>Kilograms per Square Meter, Kilograms, Decimal</desc>
  </UOM>
  <UOM>
    <code>KO</code>
    <desc>Millequivalence Caustic Potash per Gram of Product</desc>
  </UOM>
  <UOM>
    <code>KP</code>
    <desc>Kilometers Per Hour</desc>
  </UOM>
  <UOM>
    <code>KQ</code>
    <desc>Kilopascal</desc>
  </UOM>
  <UOM>
    <code>KR</code>
    <desc>Kiloroentgen</desc>
  </UOM>
  <UOM>
    <code>KS</code>
    <desc>1000 Pounds per Square Inch</desc>
  </UOM>
  <UOM>
    <code>KT</code>
    <desc>Kit</desc>
  </UOM>
  <UOM>
    <code>KU</code>
    <desc>Task</desc>
  </UOM>
  <UOM>
    <code>KV</code>
    <desc>Kelvin</desc>
  </UOM>
  <UOM>
    <code>KW</code>
    <desc>Kilograms per Millimeter</desc>
  </UOM>
  <UOM>
    <code>KX</code>
    <desc>Milliliters per Kilogram</desc>
  </UOM>
  <UOM>
    <code>K1</code>
    <desc>Kilowatt Demand</desc>
  </UOM>
  <UOM>
    <code>K2</code>
    <desc>Kilovolt Amperes Reactive Demand</desc>
  </UOM>
  <UOM>
    <code>K3</code>
    <desc>Kilovolt Amperes Reactive Hour</desc>
  </UOM>
  <UOM>
    <code>K4</code>
    <desc>Kilovolt Amperes</desc>
  </UOM>
  <UOM>
    <code>K5</code>
    <desc>Kilovolt Amperes Reactive</desc>
  </UOM>
  <UOM>
    <code>K6</code>
    <desc>Kiloliter</desc>
  </UOM>
  <UOM>
    <code>K7</code>
    <desc>Kilowatt</desc>
  </UOM>
  <UOM>
    <code>K9</code>
    <desc>Kilograms per Millimeter Squared (KG/MM2)</desc>
  </UOM>
  <UOM>
    <code>LA</code>
    <desc>Pounds Per Cubic Inch</desc>
  </UOM>
  <UOM>
    <code>LB</code>
    <desc>Pound</desc>
  </UOM>
  <UOM>
    <code>LC</code>
    <desc>Linear Centimeter</desc>
  </UOM>
  <UOM>
    <code>LE</code>
    <desc>Lite</desc>
  </UOM>
  <UOM>
    <code>LF</code>
    <desc>Linear Foot</desc>
  </UOM>
  <UOM>
    <code>LG</code>
    <desc>Long Ton</desc>
  </UOM>
  <UOM>
    <code>LH</code>
    <desc>Labor Hours</desc>
  </UOM>
  <UOM>
    <code>LI</code>
    <desc>Linear Inch</desc>
  </UOM>
  <UOM>
    <code>LJ</code>
    <desc>Large Spray</desc>
  </UOM>
  <UOM>
    <code>LK</code>
    <desc>Link</desc>
  </UOM>
  <UOM>
    <code>LL</code>
    <desc>Lifetime</desc>
  </UOM>
  <UOM>
    <code>LM</code>
    <desc>Linear Meter</desc>
  </UOM>
  <UOM>
    <code>LN</code>
    <desc>Length</desc>
  </UOM>
  <UOM>
    <code>LO</code>
    <desc>Lot</desc>
  </UOM>
  <UOM>
    <code>LP</code>
    <desc>Liquid Pounds</desc>
  </UOM>
  <UOM>
    <code>LQ</code>
    <desc>Liters Per Day</desc>
  </UOM>
  <UOM>
    <code>LR</code>
    <desc>Layer(s)</desc>
  </UOM>
  <UOM>
    <code>LS</code>
    <desc>Lump Sum</desc>
  </UOM>
  <UOM>
    <code>LT</code>
    <desc>Liter</desc>
  </UOM>
  <UOM>
    <code>LX</code>
    <desc>Linear Yards Per Pound</desc>
  </UOM>
  <UOM>
    <code>LY</code>
    <desc>Linear Yard</desc>
  </UOM>
  <UOM>
    <code>L2</code>
    <desc>Liters Per Minute</desc>
  </UOM>
  <UOM>
    <code>MA</code>
    <desc>Machine/Unit</desc>
  </UOM>
  <UOM>
    <code>MB</code>
    <desc>Millimeter-Nominal</desc>
  </UOM>
  <UOM>
    <code>MC</code>
    <desc>Microgram</desc>
  </UOM>
  <UOM>
    <code>MD</code>
    <desc>Air Dry Metric Ton</desc>
  </UOM>
  <UOM>
    <code>ME</code>
    <desc>Milligram</desc>
  </UOM>
  <UOM>
    <code>MF</code>
    <desc>Milligram per Sq. Ft. per Side</desc>
  </UOM>
  <UOM>
    <code>MG</code>
    <desc>Metric Gross Ton</desc>
  </UOM>
  <UOM>
    <code>MH</code>
    <desc>Microns (Micrometers)</desc>
  </UOM>
  <UOM>
    <code>MI</code>
    <desc>Metric</desc>
  </UOM>
  <UOM>
    <code>MJ</code>
    <desc>Minutes</desc>
  </UOM>
  <UOM>
    <code>MK</code>
    <desc>Milligrams Per Square Inch</desc>
  </UOM>
  <UOM>
    <code>ML</code>
    <desc>Milliliter</desc>
  </UOM>
  <UOM>
    <code>MN</code>
    <desc>Metric Net Ton</desc>
  </UOM>
  <UOM>
    <code>MO</code>
    <desc>Months</desc>
  </UOM>
  <UOM>
    <code>MP</code>
    <desc>Metric Ton</desc>
  </UOM>
  <UOM>
    <code>MQ</code>
    <desc>1000 Meters</desc>
  </UOM>
  <UOM>
    <code>MR</code>
    <desc>Meter</desc>
  </UOM>
  <UOM>
    <code>MS</code>
    <desc>Square Millimeter</desc>
  </UOM>
  <UOM>
    <code>MT</code>
    <desc>Metric Long Ton</desc>
  </UOM>
  <UOM>
    <code>MU</code>
    <desc>Millicurie</desc>
  </UOM>
  <UOM>
    <code>MV</code>
    <desc>Number of Mults</desc>
  </UOM>
  <UOM>
    <code>MW</code>
    <desc>Metric Ton Kilograms</desc>
  </UOM>
  <UOM>
    <code>MX</code>
    <desc>Mixed</desc>
  </UOM>
  <UOM>
    <code>MY</code>
    <desc>Millimeter-Average</desc>
  </UOM>
  <UOM>
    <code>MZ</code>
    <desc>Millimeter-minimum</desc>
  </UOM>
  <UOM>
    <code>M0</code>
    <desc>Magnetic Tapes</desc>
  </UOM>
  <UOM>
    <code>M1</code>
    <desc>Milligrams per Liter</desc>
  </UOM>
  <UOM>
    <code>M2</code>
    <desc>Millimeter-Actual</desc>
  </UOM>
  <UOM>
    <code>M3</code>
    <desc>Mat</desc>
  </UOM>
  <UOM>
    <code>M4</code>
    <desc>Monetary Value</desc>
  </UOM>
  <UOM>
    <code>M5</code>
    <desc>Microcurie</desc>
  </UOM>
  <UOM>
    <code>M6</code>
    <desc>Millibar</desc>
  </UOM>
  <UOM>
    <code>M7</code>
    <desc>Micro Inch</desc>
  </UOM>
  <UOM>
    <code>M8</code>
    <desc>Mega Pascals</desc>
  </UOM>
  <UOM>
    <code>M9</code>
    <desc>Million British Thermal Units per One Thousand Cubic</desc>
  </UOM>
  <UOM>
    <code>Feet</code>
    <desc>Code></desc>
    <code>NA</code>
    <desc>Milligrams per Kilogram</desc>
  </UOM>
  <UOM>
    <code>NB</code>
    <desc>Barge</desc>
  </UOM>
  <UOM>
    <code>NC</code>
    <desc>Car</desc>
  </UOM>
  <UOM>
    <code>ND</code>
    <desc>Net Barrels</desc>
  </UOM>
  <UOM>
    <code>NE</code>
    <desc>Net Liters</desc>
  </UOM>
  <UOM>
    <code>NF</code>
    <desc>Messages</desc>
  </UOM>
  <UOM>
    <code>NG</code>
    <desc>Net Gallons</desc>
  </UOM>
  <UOM>
    <code>NH</code>
    <desc>Message Hours</desc>
  </UOM>
  <UOM>
    <code>NI</code>
    <desc>Net Imperial Gallons</desc>
  </UOM>
  <UOM>
    <code>NJ</code>
    <desc>Number of Screens</desc>
  </UOM>
  <UOM>
    <code>NL</code>
    <desc>Load</desc>
  </UOM>
  <UOM>
    <code>NM</code>
    <desc>Nautical Mile</desc>
  </UOM>
  <UOM>
    <code>NN</code>
    <desc>Train</desc>
  </UOM>
  <UOM>
    <code>NQ</code>
    <desc>Mho</desc>
  </UOM>
  <UOM>
    <code>NR</code>
    <desc>Micro Mho</desc>
  </UOM>
  <UOM>
    <code>NS</code>
    <desc>Short Ton</desc>
  </UOM>
  <UOM>
    <code>NT</code>
    <desc>Trailer</desc>
  </UOM>
  <UOM>
    <code>NU</code>
    <desc>Newton-Meter</desc>
  </UOM>
  <UOM>
    <code>NV</code>
    <desc>Vehicle</desc>
  </UOM>
  <UOM>
    <code>NW</code>
    <desc>Newton</desc>
  </UOM>
  <UOM>
    <code>NX</code>
    <desc>Parts Per Thousand</desc>
  </UOM>
  <UOM>
    <code>NY</code>
    <desc>Pounds Per Air-Dry Metric Ton</desc>
  </UOM>
  <UOM>
    <code>N1</code>
    <desc>Pen Calories</desc>
  </UOM>
  <UOM>
    <code>N2</code>
    <desc>Number of Lines</desc>
  </UOM>
  <UOM>
    <code>N3</code>
    <desc>Print Point</desc>
  </UOM>
  <UOM>
    <code>N4</code>
    <desc>Pen Grams (Protein)</desc>
  </UOM>
  <UOM>
    <code>N6</code>
    <desc>Megahertz</desc>
  </UOM>
  <UOM>
    <code>N7</code>
    <desc>Parts</desc>
  </UOM>
  <UOM>
    <code>N9</code>
    <desc>Cartridge Needle</desc>
  </UOM>
  <UOM>
    <code>OA</code>
    <desc>Panel</desc>
  </UOM>
  <UOM>
    <code>ON</code>
    <desc>Ounces per Square Yard</desc>
  </UOM>
  <UOM>
    <code>OP</code>
    <desc>Two pack</desc>
  </UOM>
  <UOM>
    <code>OT</code>
    <desc>Overtime Hours</desc>
  </UOM>
  <UOM>
    <code>OZ</code>
    <desc>Ounce - Av</desc>
  </UOM>
  <UOM>
    <code>PA</code>
    <desc>Pail</desc>
  </UOM>
  <UOM>
    <code>PB</code>
    <desc>Pair Inches</desc>
  </UOM>
  <UOM>
    <code>PC</code>
    <desc>Piece</desc>
  </UOM>
  <UOM>
    <code>PD</code>
    <desc>Pad</desc>
  </UOM>
  <UOM>
    <code>PE</code>
    <desc>Pounds Equivalent</desc>
  </UOM>
  <UOM>
    <code>PF</code>
    <desc>Pallet (Lift)</desc>
  </UOM>
  <UOM>
    <code>PG</code>
    <desc>Pounds Gross</desc>
  </UOM>
  <UOM>
    <code>PH</code>
    <desc>Pack (PAK)</desc>
  </UOM>
  <UOM>
    <code>PI</code>
    <desc>Pitch</desc>
  </UOM>
  <UOM>
    <code>PJ</code>
    <desc>Pounds, Decimal - Pounds per Square Foot - Pound Gage</desc>
  </UOM>
  <UOM>
    <code>PK</code>
    <desc>Package</desc>
  </UOM>
  <UOM>
    <code>PL</code>
    <desc>Pallet/Unit Load</desc>
  </UOM>
  <UOM>
    <code>PM</code>
    <desc>Pounds-Percentage</desc>
  </UOM>
  <UOM>
    <code>PN</code>
    <desc>Pounds Net</desc>
  </UOM>
  <UOM>
    <code>PO</code>
    <desc>Pounds per Inch of Length</desc>
  </UOM>
  <UOM>
    <code>PP</code>
    <desc>Plate</desc>
  </UOM>
  <UOM>
    <code>PQ</code>
    <desc>Pages per Inch</desc>
  </UOM>
  <UOM>
    <code>PR</code>
    <desc>Pair</desc>
  </UOM>
  <UOM>
    <code>PS</code>
    <desc>Pounds per Sq. Inch</desc>
  </UOM>
  <UOM>
    <code>PT</code>
    <desc>Pint</desc>
  </UOM>
  <UOM>
    <code>PU</code>
    <desc>Mass Pounds</desc>
  </UOM>
  <UOM>
    <code>PV</code>
    <desc>Half Pint</desc>
  </UOM>
  <UOM>
    <code>PW</code>
    <desc>Pounds per Inch of Width</desc>
  </UOM>
  <UOM>
    <code>PX</code>
    <desc>Pint, Imperial</desc>
  </UOM>
  <UOM>
    <code>PY</code>
    <desc>Peck, Dry U.S.</desc>
  </UOM>
  <UOM>
    <code>PZ</code>
    <desc>Peck, Dry Imperial</desc>
  </UOM>
  <UOM>
    <code>P0</code>
    <desc>Pages - Electronic</desc>
  </UOM>
  <UOM>
    <code>P1</code>
    <desc>Percent</desc>
  </UOM>
  <UOM>
    <code>P2</code>
    <desc>Pounds per Foot</desc>
  </UOM>
  <UOM>
    <code>P3</code>
    <desc>Three pack</desc>
  </UOM>
  <UOM>
    <code>P4</code>
    <desc>Four-pack</desc>
  </UOM>
  <UOM>
    <code>P5</code>
    <desc>Five-pack</desc>
  </UOM>
  <UOM>
    <code>P6</code>
    <desc>Six pack</desc>
  </UOM>
  <UOM>
    <code>P7</code>
    <desc>Seven pack</desc>
  </UOM>
  <UOM>
    <code>P8</code>
    <desc>Eight-pack</desc>
  </UOM>
  <UOM>
    <code>P9</code>
    <desc>Nine pack</desc>
  </UOM>
  <UOM>
    <code>QA</code>
    <desc>Pages - Facsimile</desc>
  </UOM>
  <UOM>
    <code>QB</code>
    <desc>Pages - Hardcopy</desc>
  </UOM>
  <UOM>
    <code>QC</code>
    <desc>Channel</desc>
  </UOM>
  <UOM>
    <code>QD</code>
    <desc>Quarter Dozen</desc>
  </UOM>
  <UOM>
    <code>QE</code>
    <desc>Photographs</desc>
  </UOM>
  <UOM>
    <code>QH</code>
    <desc>Quarter Hours</desc>
  </UOM>
  <UOM>
    <code>QK</code>
    <desc>Quarter Kilogram</desc>
  </UOM>
  <UOM>
    <code>QR</code>
    <desc>Quire</desc>
  </UOM>
  <UOM>
    <code>QS</code>
    <desc>Quart, Dry U.S.</desc>
  </UOM>
  <UOM>
    <code>QT</code>
    <desc>Quart</desc>
  </UOM>
  <UOM>
    <code>QU</code>
    <desc>Quart, Imperial</desc>
  </UOM>
  <UOM>
    <code>Q1</code>
    <desc>Quarter (Time)</desc>
  </UOM>
  <UOM>
    <code>Q2</code>
    <desc>Pint U.S. Dry</desc>
  </UOM>
  <UOM>
    <code>Q3</code>
    <desc>Meal</desc>
  </UOM>
  <UOM>
    <code>Q5</code>
    <desc>Twenty-Five</desc>
  </UOM>
  <UOM>
    <code>Q6</code>
    <desc>Thirty-Six</desc>
  </UOM>
  <UOM>
    <code>Q7</code>
    <desc>Twenty-Four</desc>
  </UOM>
  <UOM>
    <code>RA</code>
    <desc>Rack</desc>
  </UOM>
  <UOM>
    <code>RB</code>
    <desc>Radian</desc>
  </UOM>
  <UOM>
    <code>RC</code>
    <desc>Rod (area) - 16.25 Square Yards</desc>
  </UOM>
  <UOM>
    <code>RD</code>
    <desc>Rod (length) - 5.5 Yards</desc>
  </UOM>
  <UOM>
    <code>RE</code>
    <desc>Reel</desc>
  </UOM>
  <UOM>
    <code>RG</code>
    <desc>Ring</desc>
  </UOM>
  <UOM>
    <code>RH</code>
    <desc>Running or Operating Hours</desc>
  </UOM>
  <UOM>
    <code>RK</code>
    <desc>Roll-Metric Measure</desc>
  </UOM>
  <UOM>
    <code>RL</code>
    <desc>Roll</desc>
  </UOM>
  <UOM>
    <code>RM</code>
    <desc>Ream</desc>
  </UOM>
  <UOM>
    <code>RN</code>
    <desc>Ream-Metric Measure</desc>
  </UOM>
  <UOM>
    <code>RO</code>
    <desc>Round</desc>
  </UOM>
  <UOM>
    <code>RP</code>
    <desc>Pounds per Ream</desc>
  </UOM>
  <UOM>
    <code>RS</code>
    <desc>Resets</desc>
  </UOM>
  <UOM>
    <code>RT</code>
    <desc>Revenue Ton Miles</desc>
  </UOM>
  <UOM>
    <code>RU</code>
    <desc>Run</desc>
  </UOM>
  <UOM>
    <code>R1</code>
    <desc>Pica</desc>
  </UOM>
  <UOM>
    <code>R2</code>
    <desc>Becquerel</desc>
  </UOM>
  <UOM>
    <code>R3</code>
    <desc>Revolutions Per Minute</desc>
  </UOM>
  <UOM>
    <code>R4</code>
    <desc>Calorie</desc>
  </UOM>
  <UOM>
    <code>R5</code>
    <desc>Thousands of Dollars</desc>
  </UOM>
  <UOM>
    <code>R6</code>
    <desc>Millions of Dollars</desc>
  </UOM>
  <UOM>
    <code>R7</code>
    <desc>Billions of Dollars</desc>
  </UOM>
  <UOM>
    <code>R8</code>
    <desc>Roentgen Equivalent in Man (REM)</desc>
  </UOM>
  <UOM>
    <code>R9</code>
    <desc>Thousand Cubic Meters</desc>
  </UOM>
  <UOM>
    <code>SA</code>
    <desc>Sandwich</desc>
  </UOM>
  <UOM>
    <code>SB</code>
    <desc>Square Mile</desc>
  </UOM>
  <UOM>
    <code>SC</code>
    <desc>Square Centimeter</desc>
  </UOM>
  <UOM>
    <code>SD</code>
    <desc>Solid Pounds</desc>
  </UOM>
  <UOM>
    <code>SE</code>
    <desc>Section</desc>
  </UOM>
  <UOM>
    <code>SF</code>
    <desc>Square Foot</desc>
  </UOM>
  <UOM>
    <code>SG</code>
    <desc>Segment</desc>
  </UOM>
  <UOM>
    <code>SH</code>
    <desc>Sheet</desc>
  </UOM>
  <UOM>
    <code>SI</code>
    <desc>Square Inch</desc>
  </UOM>
  <UOM>
    <code>SJ</code>
    <desc>Sack</desc>
  </UOM>
  <UOM>
    <code>SK</code>
    <desc>Split Tanktruck</desc>
  </UOM>
  <UOM>
    <code>SL</code>
    <desc>Sleeve</desc>
  </UOM>
  <UOM>
    <code>SM</code>
    <desc>Square Meter</desc>
  </UOM>
  <UOM>
    <code>SN</code>
    <desc>Square Rod</desc>
  </UOM>
  <UOM>
    <code>SO</code>
    <desc>Spool</desc>
  </UOM>
  <UOM>
    <code>SP</code>
    <desc>Shelf Package</desc>
  </UOM>
  <UOM>
    <code>SQ</code>
    <desc>Square</desc>
  </UOM>
  <UOM>
    <code>SR</code>
    <desc>Strip</desc>
  </UOM>
  <UOM>
    <code>SS</code>
    <desc>Sheet-Metric Measure</desc>
  </UOM>
  <UOM>
    <code>ST</code>
    <desc>Set</desc>
  </UOM>
  <UOM>
    <code>SV</code>
    <desc>Skid</desc>
  </UOM>
  <UOM>
    <code>SW</code>
    <desc>Skein</desc>
  </UOM>
  <UOM>
    <code>SX</code>
    <desc>Shipment</desc>
  </UOM>
  <UOM>
    <code>SY</code>
    <desc>Square Yard</desc>
  </UOM>
  <UOM>
    <code>SZ</code>
    <desc>Syringe</desc>
  </UOM>
  <UOM>
    <code>S1</code>
    <desc>Semester</desc>
  </UOM>
  <UOM>
    <code>S3</code>
    <desc>Square Feet per Second</desc>
  </UOM>
  <UOM>
    <code>S4</code>
    <desc>Square Meters per Second</desc>
  </UOM>
  <UOM>
    <code>S5</code>
    <desc>Sixty-fourths of an Inch</desc>
  </UOM>
  <UOM>
    <code>S6</code>
    <desc>Sessions</desc>
  </UOM>
  <UOM>
    <code>S7</code>
    <desc>Storage Units</desc>
  </UOM>
  <UOM>
    <code>S8</code>
    <desc>Standard Advertising Units (SAUs)</desc>
  </UOM>
  <UOM>
    <code>S9</code>
    <desc>Slip Sheet</desc>
  </UOM>
  <UOM>
    <code>TA</code>
    <desc>Tenth Cubic Foot</desc>
  </UOM>
  <UOM>
    <code>TB</code>
    <desc>Tube</desc>
  </UOM>
  <UOM>
    <code>TC</code>
    <desc>Truckload</desc>
  </UOM>
  <UOM>
    <code>TD</code>
    <desc>Therms</desc>
  </UOM>
  <UOM>
    <code>TE</code>
    <desc>Tote</desc>
  </UOM>
  <UOM>
    <code>TF</code>
    <desc>Ten Square Yards</desc>
  </UOM>
  <UOM>
    <code>TG</code>
    <desc>Gross Ton</desc>
  </UOM>
  <UOM>
    <code>TH</code>
    <desc>Thousand</desc>
  </UOM>
  <UOM>
    <code>TI</code>
    <desc>Thousand Square Inches</desc>
  </UOM>
  <UOM>
    <code>TJ</code>
    <desc>Thousand Sq. Centimeters</desc>
  </UOM>
  <UOM>
    <code>TK</code>
    <desc>Tank</desc>
  </UOM>
  <UOM>
    <code>TL</code>
    <desc>Thousand Feet (Linear)</desc>
  </UOM>
  <UOM>
    <code>TM</code>
    <desc>Thousand Feet (Board)</desc>
  </UOM>
  <UOM>
    <code>TN</code>
    <desc>Net Ton (2,000 LB).</desc>
  </UOM>
  <UOM>
    <code>TO</code>
    <desc>Troy Ounce</desc>
  </UOM>
  <UOM>
    <code>TP</code>
    <desc>Ten-pack</desc>
  </UOM>
  <UOM>
    <code>TQ</code>
    <desc>Thousand Feet</desc>
  </UOM>
  <UOM>
    <code>TR</code>
    <desc>Ten Square Feet</desc>
  </UOM>
  <UOM>
    <code>TS</code>
    <desc>Thousand Square Feet</desc>
  </UOM>
  <UOM>
    <code>TT</code>
    <desc>Thousand Linear Meters</desc>
  </UOM>
  <UOM>
    <code>TU</code>
    <desc>Thousand Linear Yards</desc>
  </UOM>
  <UOM>
    <code>TV</code>
    <desc>Thousand Kilograms</desc>
  </UOM>
  <UOM>
    <code>TW</code>
    <desc>Thousand Sheets</desc>
  </UOM>
  <UOM>
    <code>TX</code>
    <desc>Troy Pound</desc>
  </UOM>
  <UOM>
    <code>TY</code>
    <desc>Tray</desc>
  </UOM>
  <UOM>
    <code>TZ</code>
    <desc>Thousand Cubic Feet</desc>
  </UOM>
  <UOM>
    <code>T0</code>
    <desc>Telecommunications Lines in Service</desc>
  </UOM>
  <UOM>
    <code>T1</code>
    <desc>Thousand pounds gross</desc>
  </UOM>
  <UOM>
    <code>T2</code>
    <desc>Thousandths of an Inch</desc>
  </UOM>
  <UOM>
    <code>T3</code>
    <desc>Thousand Pieces</desc>
  </UOM>
  <UOM>
    <code>T4</code>
    <desc>Thousand Bags</desc>
  </UOM>
  <UOM>
    <code>T5</code>
    <desc>Thousand Casings</desc>
  </UOM>
  <UOM>
    <code>T6</code>
    <desc>Thousand Gallons</desc>
  </UOM>
  <UOM>
    <code>T7</code>
    <desc>Thousand Impressions</desc>
  </UOM>
  <UOM>
    <code>T8</code>
    <desc>Thousand Linear Inches</desc>
  </UOM>
  <UOM>
    <code>T9</code>
    <desc>Thousand Kilowatt Hours</desc>
  </UOM>
  <UOM>
    <code>UA</code>
    <desc>Torr</desc>
  </UOM>
  <UOM>
    <code>UB</code>
    <desc>Telecommunications Lines in Service - Average</desc>
  </UOM>
  <UOM>
    <code>UC</code>
    <desc>Telecommunications Ports</desc>
  </UOM>
  <UOM>
    <code>UD</code>
    <desc>Tenth Minutes</desc>
  </UOM>
  <UOM>
    <code>UE</code>
    <desc>Tenth Hours</desc>
  </UOM>
  <UOM>
    <code>UF</code>
    <desc>Usage per Telecommunications Line - Average</desc>
  </UOM>
  <UOM>
    <code>UH</code>
    <desc>Ten Thousand Yards</desc>
  </UOM>
  <UOM>
    <code>UL</code>
    <desc>Unitless</desc>
  </UOM>
  <UOM>
    <code>UM</code>
    <desc>Million Units</desc>
  </UOM>
  <UOM>
    <code>UN</code>
    <desc>Unit</desc>
  </UOM>
  <UOM>
    <code>UP</code>
    <desc>Troche</desc>
  </UOM>
  <UOM>
    <code>UR</code>
    <desc>Application</desc>
  </UOM>
  <UOM>
    <code>US</code>
    <desc>Dosage Form</desc>
  </UOM>
  <UOM>
    <code>UT</code>
    <desc>Inhalation</desc>
  </UOM>
  <UOM>
    <code>UU</code>
    <desc>Lozenge</desc>
  </UOM>
  <UOM>
    <code>UV</code>
    <desc>Percent Topical Only</desc>
  </UOM>
  <UOM>
    <code>UW</code>
    <desc>Milliequivalent</desc>
  </UOM>
  <UOM>
    <code>UX</code>
    <desc>Dram (Minim)</desc>
  </UOM>
  <UOM>
    <code>UY</code>
    <desc>Fifty Square Feet</desc>
  </UOM>
  <UOM>
    <code>UZ</code>
    <desc>Fifty Count</desc>
  </UOM>
  <UOM>
    <code>U1</code>
    <desc>Treatments</desc>
  </UOM>
  <UOM>
    <code>U2</code>
    <desc>Tablet</desc>
  </UOM>
  <UOM>
    <code>U3</code>
    <desc>Ten</desc>
  </UOM>
  <UOM>
    <code>U5</code>
    <desc>Two Hundred Fifty</desc>
  </UOM>
  <UOM>
    <code>VA</code>
    <desc>Volt-ampere per Kilogram</desc>
  </UOM>
  <UOM>
    <code>VC</code>
    <desc>Five Hundred</desc>
  </UOM>
  <UOM>
    <code>VI</code>
    <desc>Vial</desc>
  </UOM>
  <UOM>
    <code>VP</code>
    <desc>Percent Volume</desc>
  </UOM>
  <UOM>
    <code>VR</code>
    <desc>Volt-ampere-reactive</desc>
  </UOM>
  <UOM>
    <code>VS</code>
    <desc>Visit</desc>
  </UOM>
  <UOM>
    <code>V1</code>
    <desc>Flat</desc>
  </UOM>
  <UOM>
    <code>V2</code>
    <desc>Pouch</desc>
  </UOM>
  <UOM>
    <code>WA</code>
    <desc>Watts per Kilogram</desc>
  </UOM>
  <UOM>
    <code>WB</code>
    <desc>Wet Pound</desc>
  </UOM>
  <UOM>
    <code>WD</code>
    <desc>Work Days</desc>
  </UOM>
  <UOM>
    <code>WE</code>
    <desc>Wet Ton</desc>
  </UOM>
  <UOM>
    <code>WG</code>
    <desc>Wine Gallon</desc>
  </UOM>
  <UOM>
    <code>WH</code>
    <desc>Wheel</desc>
  </UOM>
  <UOM>
    <code>WI</code>
    <desc>Weight per Square Inch</desc>
  </UOM>
  <UOM>
    <code>WK</code>
    <desc>Week</desc>
  </UOM>
  <UOM>
    <code>WM</code>
    <desc>Working Months</desc>
  </UOM>
  <UOM>
    <code>WP</code>
    <desc>Pennyweight</desc>
  </UOM>
  <UOM>
    <code>WR</code>
    <desc>Wrap</desc>
  </UOM>
  <UOM>
    <code>WW</code>
    <desc>Milliliters of Water</desc>
  </UOM>
  <UOM>
    <code>W2</code>
    <desc>Wet Kilo</desc>
  </UOM>
  <UOM>
    <code>XP</code>
    <desc>Base Box per Pound</desc>
  </UOM>
  <UOM>
    <code>X1</code>
    <desc>Chains (Land Survey)</desc>
  </UOM>
  <UOM>
    <code>X2</code>
    <desc>Bunch</desc>
  </UOM>
  <UOM>
    <code>X3</code>
    <desc>Clove</desc>
  </UOM>
  <UOM>
    <code>X4</code>
    <desc>Drop</desc>
  </UOM>
  <UOM>
    <code>X5</code>
    <desc>Head</desc>
  </UOM>
  <UOM>
    <code>X6</code>
    <desc>Heart</desc>
  </UOM>
  <UOM>
    <code>X7</code>
    <desc>Leaf</desc>
  </UOM>
  <UOM>
    <code>X8</code>
    <desc>Loaf</desc>
  </UOM>
  <UOM>
    <code>X9</code>
    <desc>Portion</desc>
  </UOM>
  <UOM>
    <code>YD</code>
    <desc>Yard</desc>
  </UOM>
  <UOM>
    <code>YL</code>
    <desc>100 Lineal Yards</desc>
  </UOM>
  <UOM>
    <code>YR</code>
    <desc>Years</desc>
  </UOM>
  <UOM>
    <code>YT</code>
    <desc>Ten Yards</desc>
  </UOM>
  <UOM>
    <code>Y1</code>
    <desc>Slice</desc>
  </UOM>
  <UOM>
    <code>Y2</code>
    <desc>Tablespoon</desc>
  </UOM>
  <UOM>
    <code>Y3</code>
    <desc>Teaspoon</desc>
  </UOM>
  <UOM>
    <code>Y4</code>
    <desc>Tub</desc>
  </UOM>
  <UOM>
    <code>ZA</code>
    <desc>Bimonthly</desc>
  </UOM>
  <UOM>
    <code>ZB</code>
    <desc>Biweekly</desc>
  </UOM>
  <UOM>
    <code>ZP</code>
    <desc>Page</desc>
  </UOM>
  <UOM>
    <code>ZZ</code>
    <desc>Mutually Defined</desc>
  </UOM>
  <UOM>
    <code>Z1</code>
    <desc>Lift Van</desc>
  </UOM>
  <UOM>
    <code>Z2</code>
    <desc>Chest</desc>
  </UOM>
  <UOM>
    <code>Z3</code>
    <desc>Cask</desc>
  </UOM>
  <UOM>
    <code>Z4</code>
    <desc>Hogshead</desc>
  </UOM>
  <UOM>
    <code>Z5</code>
    <desc>Lug</desc>
  </UOM>
  <UOM>
    <code>Z6</code>
    <desc>Conference Points</desc>
  </UOM>
  <UOM>
    <code>Z8</code>
    <desc>Newspaper Agate Line</desc>
  </UOM>
  <UOM>
    <code>1</code>
    <desc>ctual Pounds</desc>
  </UOM>
  <UOM>
    <code>2</code>
    <desc>tatute Mile</desc>
  </UOM>
  <UOM>
    <code>3</code>
    <desc>econds</desc>
  </UOM>
  <UOM>
    <code>4</code>
    <desc>mall Spray</desc>
  </UOM>
  <UOM>
    <code>5</code>
    <desc>ifts</desc>
  </UOM>
  <UOM>
    <code>6</code>
    <desc>igits</desc>
  </UOM>
  <UOM>
    <code>7</code>
    <desc>trand</desc>
  </UOM>
  <UOM>
    <code>8</code>
    <desc>eat Lots</desc>
  </UOM>
  <UOM>
    <code>9</code>
    <desc>ire</desc>
  </UOM>
  <UOM>
    <code>1A</code>
    <desc>Car Mile</desc>
  </UOM>
  <UOM>
    <code>1B</code>
    <desc>Car Count</desc>
  </UOM>
  <UOM>
    <code>1C</code>
    <desc>Locomotive Count</desc>
  </UOM>
  <UOM>
    <code>1D</code>
    <desc>Caboose Count</desc>
  </UOM>
  <UOM>
    <code>1E</code>
    <desc>Empty Car</desc>
  </UOM>
  <UOM>
    <code>1F</code>
    <desc>Train Mile</desc>
  </UOM>
  <UOM>
    <code>1G</code>
    <desc>Fuel Usage (Gallons)</desc>
  </UOM>
  <UOM>
    <code>1H</code>
    <desc>Caboose Mile</desc>
  </UOM>
  <UOM>
    <code>1I</code>
    <desc>Fixed Rate</desc>
  </UOM>
  <UOM>
    <code>1J</code>
    <desc>Ton Miles</desc>
  </UOM>
  <UOM>
    <code>1K</code>
    <desc>Locomotive Mile</desc>
  </UOM>
  <UOM>
    <code>1L</code>
    <desc>Total Car Count</desc>
  </UOM>
  <UOM>
    <code>1M</code>
    <desc>Total Car Mile</desc>
  </UOM>
  <UOM>
    <code>1N</code>
    <desc>Count</desc>
  </UOM>
  <UOM>
    <code>1O</code>
    <desc>Season</desc>
  </UOM>
  <UOM>
    <code>1P</code>
    <desc>Tank Car</desc>
  </UOM>
  <UOM>
    <code>1Q</code>
    <desc>Frames</desc>
  </UOM>
  <UOM>
    <code>1R</code>
    <desc>Transactions</desc>
  </UOM>
  <UOM>
    <code>1X</code>
    <desc>Quarter Mile</desc>
  </UOM>
  <UOM>
    <code>10</code>
    <desc>Group</desc>
  </UOM>
  <UOM>
    <code>11</code>
    <desc>Outfit</desc>
  </UOM>
  <UOM>
    <code>12</code>
    <desc>Packet</desc>
  </UOM>
  <UOM>
    <code>13</code>
    <desc>Ration</desc>
  </UOM>
  <UOM>
    <code>14</code>
    <desc>Shot</desc>
  </UOM>
  <UOM>
    <code>15</code>
    <desc>Stick</desc>
  </UOM>
  <UOM>
    <code>16</code>
    <desc>115 Kilogram Drum</desc>
  </UOM>
  <UOM>
    <code>17</code>
    <desc>100 Pound Drum</desc>
  </UOM>
  <UOM>
    <code>18</code>
    <desc>55 Gallon Drum</desc>
  </UOM>
  <UOM>
    <code>19</code>
    <desc>Tank Truck</desc>
  </UOM>
  <UOM>
    <code>2A</code>
    <desc>Radians Per Second</desc>
  </UOM>
  <UOM>
    <code>2B</code>
    <desc>Radians Per Second Squared</desc>
  </UOM>
  <UOM>
    <code>2C</code>
    <desc>Roentgen</desc>
  </UOM>
  <UOM>
    <code>2F</code>
    <desc>Volts Per Meter</desc>
  </UOM>
  <UOM>
    <code>2G</code>
    <desc>Volts (Alternating Current)</desc>
  </UOM>
  <UOM>
    <code>2H</code>
    <desc>Volts (Direct Current)</desc>
  </UOM>
  <UOM>
    <code>2I</code>
    <desc>British Thermal Units (BTUs) Per Hour</desc>
  </UOM>
  <UOM>
    <code>2K</code>
    <desc>Cubic Feet Per Hour</desc>
  </UOM>
  <UOM>
    <code>2L</code>
    <desc>Cubic Feet Per Minute</desc>
  </UOM>
  <UOM>
    <code>2M</code>
    <desc>Centimeters Per Second</desc>
  </UOM>
  <UOM>
    <code>2N</code>
    <desc>Decibels</desc>
  </UOM>
  <UOM>
    <code>2P</code>
    <desc>Kilobyte</desc>
  </UOM>
  <UOM>
    <code>2Q</code>
    <desc>Kilobecquerel</desc>
  </UOM>
  <UOM>
    <code>2R</code>
    <desc>Kilocurie</desc>
  </UOM>
  <UOM>
    <code>2U</code>
    <desc>Megagram</desc>
  </UOM>
  <UOM>
    <code>2V</code>
    <desc>Megagrams Per Hour</desc>
  </UOM>
  <UOM>
    <code>2W</code>
    <desc>Bin</desc>
  </UOM>
  <UOM>
    <code>2X</code>
    <desc>Meters Per Minute</desc>
  </UOM>
  <UOM>
    <code>2Y</code>
    <desc>Milliroentgen</desc>
  </UOM>
  <UOM>
    <code>2Z</code>
    <desc>Millivolts</desc>
  </UOM>
  <UOM>
    <code>20</code>
    <desc>20 Foot Container</desc>
  </UOM>
  <UOM>
    <code>21</code>
    <desc>40 Foot Container</desc>
  </UOM>
  <UOM>
    <code>22</code>
    <desc>Deciliter per Gram</desc>
  </UOM>
  <UOM>
    <code>23</code>
    <desc>Grams per Cubic Centimeter</desc>
  </UOM>
  <UOM>
    <code>24</code>
    <desc>Theoretical Pounds</desc>
  </UOM>
  <UOM>
    <code>25</code>
    <desc>Grams per Square Centimeter</desc>
  </UOM>
  <UOM>
    <code>26</code>
    <desc>Actual Tons</desc>
  </UOM>
  <UOM>
    <code>27</code>
    <desc>Theoretical Tons</desc>
  </UOM>
  <UOM>
    <code>28</code>
    <desc>Kilograms per Square Meter</desc>
  </UOM>
  <UOM>
    <code>29</code>
    <desc>Pounds per 1000 Square Feet</desc>
  </UOM>
  <UOM>
    <code>3B</code>
    <desc>Megajoule</desc>
  </UOM>
  <UOM>
    <code>3C</code>
    <desc>Manmonth</desc>
  </UOM>
  <UOM>
    <code>3E</code>
    <desc>Pounds Per Pound of Product</desc>
  </UOM>
  <UOM>
    <code>3F</code>
    <desc>Kilograms Per Liter of Product</desc>
  </UOM>
  <UOM>
    <code>3G</code>
    <desc>Pounds Per Piece of Product</desc>
  </UOM>
  <UOM>
    <code>3H</code>
    <desc>Kilograms Per Kilogram of Product</desc>
  </UOM>
  <UOM>
    <code>3I</code>
    <desc>Kilograms Per Piece of Product</desc>
  </UOM>
  <UOM>
    <code>30</code>
    <desc>Horsepower Days per Air Dry Metric Tons</desc>
  </UOM>
  <UOM>
    <code>31</code>
    <desc>Catchweight</desc>
  </UOM>
  <UOM>
    <code>32</code>
    <desc>Kilograms per Air Dry Metric Tons</desc>
  </UOM>
  <UOM>
    <code>33</code>
    <desc>Kilopascal Square Meters per Gram</desc>
  </UOM>
  <UOM>
    <code>34</code>
    <desc>Kilopascals per Millimeter</desc>
  </UOM>
  <UOM>
    <code>35</code>
    <desc>Milliliters per Square Centimeter Second</desc>
  </UOM>
  <UOM>
    <code>36</code>
    <desc>Cubic Feet per Minute per Square Foot</desc>
  </UOM>
  <UOM>
    <code>37</code>
    <desc>Ounces per Square Foot</desc>
  </UOM>
  <UOM>
    <code>38</code>
    <desc>Ounces per Square Foot per 0.01 Inch</desc>
  </UOM>
  <UOM>
    <code>39</code>
    <desc>Basis Points</desc>
  </UOM>
  <UOM>
    <code>4A</code>
    <desc>Bobbin</desc>
  </UOM>
  <UOM>
    <code>4B</code>
    <desc>Cap</desc>
  </UOM>
  <UOM>
    <code>4C</code>
    <desc>Centistokes</desc>
  </UOM>
  <UOM>
    <code>4D</code>
    <desc>Curie</desc>
  </UOM>
  <UOM>
    <code>4E</code>
    <desc>20-Pack</desc>
  </UOM>
  <UOM>
    <code>4F</code>
    <desc>100-Pack</desc>
  </UOM>
  <UOM>
    <code>4G</code>
    <desc>Microliter</desc>
  </UOM>
  <UOM>
    <code>4H</code>
    <desc>Micrometer</desc>
  </UOM>
  <UOM>
    <code>4I</code>
    <desc>Meters Per Second</desc>
  </UOM>
  <UOM>
    <code>4J</code>
    <desc>Meters Per Second Per Second</desc>
  </UOM>
  <UOM>
    <code>4K</code>
    <desc>Milliamperes</desc>
  </UOM>
  <UOM>
    <code>4L</code>
    <desc>Megabyte</desc>
  </UOM>
  <UOM>
    <code>4M</code>
    <desc>Milligrams Per Hour</desc>
  </UOM>
  <UOM>
    <code>4N</code>
    <desc>Megabecquerel</desc>
  </UOM>
  <UOM>
    <code>4P</code>
    <desc>Newtons Per Meter</desc>
  </UOM>
  <UOM>
    <code>4Q</code>
    <desc>Ounce Inch</desc>
  </UOM>
  <UOM>
    <code>4R</code>
    <desc>Ounce Foot</desc>
  </UOM>
  <UOM>
    <code>4S</code>
    <desc>Pascal</desc>
  </UOM>
  <UOM>
    <code>4T</code>
    <desc>Picofarad</desc>
  </UOM>
  <UOM>
    <code>4U</code>
    <desc>Pounds Per Hour</desc>
  </UOM>
  <UOM>
    <code>4V</code>
    <desc>Cubic Meter Per Hour</desc>
  </UOM>
  <UOM>
    <code>4W</code>
    <desc>Ton Per Hour</desc>
  </UOM>
  <UOM>
    <code>4X</code>
    <desc>Kiloliter Per Hour</desc>
  </UOM>
  <UOM>
    <code>40</code>
    <desc>Milliliter per Second</desc>
  </UOM>
  <UOM>
    <code>41</code>
    <desc>Milliliter per Minute</desc>
  </UOM>
  <UOM>
    <code>43</code>
    <desc>Super Bulk Bag</desc>
  </UOM>
  <UOM>
    <code>44</code>
    <desc>500 Kilogram Bulk Bag</desc>
  </UOM>
  <UOM>
    <code>45</code>
    <desc>300 Kilogram Bulk Bag</desc>
  </UOM>
  <UOM>
    <code>46</code>
    <desc>25 Kilogram Bulk Bag</desc>
  </UOM>
  <UOM>
    <code>47</code>
    <desc>50 Pound Bag</desc>
  </UOM>
  <UOM>
    <code>48</code>
    <desc>Bulk Car Load</desc>
  </UOM>
  <UOM>
    <code>5A</code>
    <desc>Barrels per Minute</desc>
  </UOM>
  <UOM>
    <code>5B</code>
    <desc>Batch</desc>
  </UOM>
  <UOM>
    <code>5C</code>
    <desc>Gallons per Thousand</desc>
  </UOM>
  <UOM>
    <code>5E</code>
    <desc>MMSCF/Day</desc>
  </UOM>
  <UOM>
    <code>5F</code>
    <desc>Pounds per Thousand</desc>
  </UOM>
  <UOM>
    <code>5G</code>
    <desc>Pump</desc>
  </UOM>
  <UOM>
    <code>5H</code>
    <desc>Stage</desc>
  </UOM>
  <UOM>
    <code>5I</code>
    <desc>Standard Cubic Foot</desc>
  </UOM>
  <UOM>
    <code>5J</code>
    <desc>Hydraulic Horse Power</desc>
  </UOM>
  <UOM>
    <code>5K</code>
    <desc>Count per Minute</desc>
  </UOM>
  <UOM>
    <code>5P</code>
    <desc>Seismic Level</desc>
  </UOM>
  <UOM>
    <code>5Q</code>
    <desc>Seismic Line</desc>
  </UOM>
  <UOM>
    <code>50</code>
    <desc>Actual Kilograms</desc>
  </UOM>
  <UOM>
    <code>51</code>
    <desc>Actual Tonnes</desc>
  </UOM>
  <UOM>
    <code>52</code>
    <desc>Credits</desc>
  </UOM>
  <UOM>
    <code>53</code>
    <desc>Theoretical Kilograms</desc>
  </UOM>
  <UOM>
    <code>54</code>
    <desc>Theoretical Tonnes</desc>
  </UOM>
  <UOM>
    <code>56</code>
    <desc>Sitas</desc>
  </UOM>
  <UOM>
    <code>57</code>
    <desc>Mesh</desc>
  </UOM>
  <UOM>
    <code>58</code>
    <desc>Net Kilograms</desc>
  </UOM>
  <UOM>
    <code>59</code>
    <desc>Parts Per Million</desc>
  </UOM>
  <UOM>
    <code>60</code>
    <desc>Percent Weight</desc>
  </UOM>
  <UOM>
    <code>61</code>
    <desc>Parts Per Billion</desc>
  </UOM>
  <UOM>
    <code>62</code>
    <desc>Percent Per 1000 Hours</desc>
  </UOM>
  <UOM>
    <code>63</code>
    <desc>Failure Rate In Time</desc>
  </UOM>
  <UOM>
    <code>64</code>
    <desc>Pounds Per Square Inch Gauge</desc>
  </UOM>
  <UOM>
    <code>65</code>
    <desc>Coulomb</desc>
  </UOM>
  <UOM>
    <code>66</code>
    <desc>Oersteds</desc>
  </UOM>
  <UOM>
    <code>67</code>
    <desc>Siemens</desc>
  </UOM>
  <UOM>
    <code>68</code>
    <desc>Ampere</desc>
  </UOM>
  <UOM>
    <code>69</code>
    <desc>Test Specific Scale</desc>
  </UOM>
  <UOM>
    <code>70</code>
    <desc>Volt</desc>
  </UOM>
  <UOM>
    <code>71</code>
    <desc>Volt-Ampere Per Pound</desc>
  </UOM>
  <UOM>
    <code>72</code>
    <desc>Watts Per Pound</desc>
  </UOM>
  <UOM>
    <code>73</code>
    <desc>Ampere Turn Per Centimeter</desc>
  </UOM>
  <UOM>
    <code>74</code>
    <desc>Milli Pascals</desc>
  </UOM>
  <UOM>
    <code>76</code>
    <desc>Gauss</desc>
  </UOM>
  <UOM>
    <code>78</code>
    <desc>Kilogauss</desc>
  </UOM>
  <UOM>
    <code>79</code>
    <desc>Electron Volt</desc>
  </UOM>
  <UOM>
    <code>8C</code>
    <desc>Cord</desc>
  </UOM>
  <UOM>
    <code>8D</code>
    <desc>Duty</desc>
  </UOM>
  <UOM>
    <code>8P</code>
    <desc>Project</desc>
  </UOM>
  <UOM>
    <code>8R</code>
    <desc>Program</desc>
  </UOM>
  <UOM>
    <code>8S</code>
    <desc>Session</desc>
  </UOM>
  <UOM>
    <code>8U</code>
    <desc>Square Kilometer</desc>
  </UOM>
  <UOM>
    <code>80</code>
    <desc>Pounds Per Square Inch Absolute</desc>
  </UOM>
  <UOM>
    <code>81</code>
    <desc>Henry</desc>
  </UOM>
  <UOM>
    <code>82</code>
    <desc>Ohm</desc>
  </UOM>
  <UOM>
    <code>83</code>
    <desc>Farad</desc>
  </UOM>
  <UOM>
    <code>84</code>
    <desc>Kilo Pounds Per Square Inch (KSI)</desc>
  </UOM>
  <UOM>
    <code>85</code>
    <desc>Foot Pounds</desc>
  </UOM>
  <UOM>
    <code>86</code>
    <desc>Joules</desc>
  </UOM>
  <UOM>
    <code>87</code>
    <desc>Pounds per Cubic Foot</desc>
  </UOM>
  <UOM>
    <code>89</code>
    <desc>Poise</desc>
  </UOM>
  <UOM>
    <code>90</code>
    <desc>Saybold Universal Second</desc>
  </UOM>
  <UOM>
    <code>91</code>
    <desc>Stokes</desc>
  </UOM>
  <UOM>
    <code>92</code>
    <desc>Calories per Cubic Centimeter</desc>
  </UOM>
  <UOM>
    <code>93</code>
    <desc>Calories per Gram</desc>
  </UOM>
  <UOM>
    <code>94</code>
    <desc>Curl Units</desc>
  </UOM>
  <UOM>
    <code>95</code>
    <desc>20,000 Gallon Tankcar</desc>
  </UOM>
  <UOM>
    <code>96</code>
    <desc>10,000 Gallon Tankcar</desc>
  </UOM>
  <UOM>
    <code>97</code>
    <desc>10 Kilogram Drum</desc>
  </UOM>
  <UOM>
    <code>98</code>
    <desc>15 Kilogram Drum</desc>
  </UOM>
  <UOM>
    <code>99</code>
    <desc>Watt</desc>
  </UOM>
</Data>
`
