export function processData(masterList) {
  let map = {};
  let zones = [];
  let crags = [];
  let areas = [];
  let blocs = [];

  const data = masterList;

  data.forEach((row) => {
    //zone
    if (!map.hasOwnProperty(row.Zone)) {
      map[row.Zone] = {};
      zones.push(row.Zone);
    }

    //crag
    if (!map[row.Zone].hasOwnProperty(row.Crag)) {
      map[row.Zone][row.Crag] = {};
      crags.push(row.Crag);
    }

    //area
    if (!map[row.Zone][row.Crag].hasOwnProperty(row.Area)) {
      map[row.Zone][row.Crag][row.Area] = {};
      areas.push(row.Area);
    }

    //boulder
    if (!map[row.Zone][row.Crag][row.Area].hasOwnProperty(row.Bloc)) {
      map[row.Zone][row.Crag][row.Area][row.Bloc] = { lines: [] };
      blocs.push(row.Bloc);
    }

    map[row.Zone][row.Crag][row.Area][row.Bloc].lines.push(row);
  });

  return map;
}