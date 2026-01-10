// DataProcessor.js
export function processData(masterList) {
  let map = {};

  masterList.forEach((row) => {
    // 1. Setup nested objects
    if (!map[row.Zone]) map[row.Zone] = {};
    if (!map[row.Zone][row.Crag]) map[row.Zone][row.Crag] = {};
    if (!map[row.Zone][row.Crag][row.Area]) map[row.Zone][row.Crag][row.Area] = {};
    if (!map[row.Zone][row.Crag][row.Area][row.Bloc]) {
      map[row.Zone][row.Crag][row.Area][row.Bloc] = {};
    }

    // 2. DEFINE THE VARIABLE HERE
    const faceName = row["Bloc Face"];

    // 3. Use the variable to initialize the Face object
    if (!map[row.Zone][row.Crag][row.Area][row.Bloc][faceName]) {
      map[row.Zone][row.Crag][row.Area][row.Bloc][faceName] = { lines: [] };
    }

    // 4. Push the row into the lines array
    map[row.Zone][row.Crag][row.Area][row.Bloc][faceName].lines.push(row);
  });

  return map;
}
