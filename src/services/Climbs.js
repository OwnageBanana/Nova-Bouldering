const BASE_URL = __API_URL + '/climbs' // Change to your full API URL if needed

async function GetAllClimbs() {
  const response = await fetch(`${BASE_URL}`)
  if (!response.ok) throw new Error('Failed to fetch climbs')
  return await response.json()
}

async function GetClimb(id) {
  const response = await fetch(`${BASE_URL}/${id}`)
  if (!response.ok) throw new Error('Failed to fetch climb')
  return await response.json()
}

async function CreateClimb(climb) {
  const response = await fetch(BASE_URL, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(climb),
  })
  if (!response.ok) throw new Error('Failed to create climb')
  return await response.json()
}

async function UpdateClimb(climb) {
  // Using POST based on your mux definition (usually this is PUT or PATCH)
  const response = await fetch(`${BASE_URL}/${climb.id}`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(climb),
  })
  if (!response.ok) throw new Error('Failed to update climb')
  return await response.json()
}

async function DeleteClimb(id) {
  const response = await fetch(`${BASE_URL}/${id}`, {
    method: 'DELETE',
  })
  if (!response.ok) throw new Error('Failed to delete climb')
  return true
}

export default { GetAllClimbs, GetClimb, CreateClimb, UpdateClimb, DeleteClimb }
