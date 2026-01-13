<template>
  <div class="container">
    <h2>API Tester</h2>
    auth lasts for 24 hours in a cookie. <br />
    <div class="section">
      <input v-model="authKey" type="password" placeholder="Auth Key" />
      <button @click="auth">Get Write Access</button>
    </div>

    <div class="section">
      <button @click="getAllClimbs">Get All Climbs</button>
    </div>

    <div class="section">
      <input v-model="newClimb.name" placeholder="Climb Name" />
      <input v-model="newClimb.grade" placeholder="Grade" />
      <button @click="createClimb">Create Climb</button>
    </div>

    <div class="section">
      <input v-model="updateData.id" placeholder="ID to Update" />
      <input v-model="updateData.boulder_id" placeholder="New boulder ref" />
      <input v-model="updateData.name" placeholder="New Name" />
      <input v-model="updateData.description" placeholder="New description" />
      <input
        v-model="updateData.face"
        placeholder="New Face Direction (north,south,east,west,'')"
      />
      <input v-model="updateData.grade" placeholder="New grade" />
      <input v-model="updateData.line" placeholder="New line" />
      <input v-model="updateData.metadata" placeholder="New metadata" />
      <div>
        <button @click="getClimb">GetClimb from id</button>
        <button @click="updateClimb">Update Climb (POST)</button>
      </div>
    </div>

    <div class="section">
      <input v-model="deleteId" placeholder="ID to Delete" />
      <button @click="deleteClimb">Delete Climb</button>
    </div>

    <div class="output">
      <h3>Response Log:</h3>
      <pre>{{ log }}</pre>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, toRaw } from 'vue'
import ClimbsService from '@services/Climbs.js'
import AuthService from '@services/Auth.js'

const BASE_URL = '/climbs'
const log = ref('Waiting for action...')
const deleteId = ref('')
const authKey = ref('')

// Form Models
const newClimb = ref({ name: '', grade: '' })
const updateData = ref({ id: 1 })

// --- API Functions ---
async function auth() {
  if (!authKey.value) return alert('Key required')
  try {
    const data = await AuthService.Auth(authKey.value)
    console.log(data.body.toString())
    log.value = JSON.stringify(data, null, 2)
    authKey.value = ''
  } catch (err) {
    console.log(err)
    log.value = `Error: ${err.message}`
  }
}

async function getAllClimbs() {
  try {
    const data = await ClimbsService.GetAllClimbs()
    log.value = JSON.stringify(data, null, 2)
  } catch (err) {
    log.value = `Error: ${err.message}`
  }
}

async function getClimb() {
  if (!updateData.value.id) return alert('ID required')
  try {
    const data = await ClimbsService.GetClimb(updateData.value.id)
    console.log(data)

    updateData.value = data
    updateData.value.line = JSON.stringify(data.line)
    updateData.value.metadata = JSON.stringify(data.metadata)
    log.value = JSON.stringify(data, null, 2)
  } catch (err) {
    console.log(err)
    log.value = `Error: ${err.message}`
  }
}

async function createClimb() {
  try {
    const data = await ClimbsService.CreateClimb(newClimb.value)
    updateData.value = data
    log.value = `Created:\n${JSON.stringify(data, null, 2)}`
  } catch (err) {
    console.log(err)
    log.value = `Error: ${err.message}`
  }
}

async function updateClimb() {
  if (!updateData.value.id) return alert('ID required')

  try {
    updateData.value.line = JSON.parse(updateData.value.line)
    updateData.value.metadata = JSON.parse(updateData.value.metadata)
    const data = await ClimbsService.UpdateClimb(updateData.value)
    updateData.value.line = JSON.stringify(data.line)
    updateData.value.metadata = JSON.stringify(data.metadata)
    log.value = `Updated:\n${JSON.stringify(data, null, 2)}`
  } catch (err) {
    console.log(err)
    log.value = `Error: ${err.message}`
  }
}

async function deleteClimb() {
  if (!deleteId.value) return alert('ID required')

  try {
    // Fixed typo from "DELET" to "DELETE"
    const res = ClimbsService.DeleteClimb(deleteId.value)
    if (res.ok) {
      log.value = `Deleted ID: ${deleteId.value}`
    } else {
      log.value = 'Failed to delete'
    }
  } catch (err) {
    console.log(err)
    log.value = `Error: ${err.message}`
  }
}
</script>

<style scoped>
.container {
  font-family: sans-serif;
  max-width: 600px;
  margin: 2rem auto;
}
.section {
  margin-bottom: 1rem;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
}
input {
  margin-right: 0.5rem;
  padding: 5px;
}
button {
  padding: 5px 10px;
  cursor: pointer;
}
.output {
  background: #f4f4f4;
  padding: 1rem;
  margin-top: 1rem;
}
</style>
