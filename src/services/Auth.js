const BASE_URL = __API_URL
async function Auth (key) {
  const response = await fetch(`${BASE_URL}/AuthWriteAccess`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ key: key }),
  })
  if (!response.ok) throw new Error('Failed to authenticate')
  return await response
}

export default { Auth }
