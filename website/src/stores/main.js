/*
 * @Author: SpenserCai
 * @Date: 2023-10-01 10:22:20
 * @version: 
 * @LastEditors: SpenserCai
 * @LastEditTime: 2023-10-04 00:36:13
 * @Description: file content
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { userinfo } from '@/api/account'
import axios from 'axios'

export const useMainStore = defineStore('main', () => {
  const userName = ref('User')
  const userEmail = ref('doe.doe.doe@example.com')
  const created = ref("")
  const stableConfig = ref({})

  const userAvatar = ref('')

  userinfo().then((res) => {
    userName.value = res.data.user.username
    // 如果res.data.user.avatar不为空，就用res.data.user.avatar，否则用默认的
    if (res.data.user.avatar != "") {
      userAvatar.value = res.data.user.avatar
    } else {
      userAvatar.value = `https://api.dicebear.com/7.x/avataaars/svg?seed=${userEmail.value.replace(
        /[^a-z0-9]+/gi,
        '-'
      )}`
    }
    created.value = res.data.user.created
    stableConfig.value = res.data.user.stable_config
  })

  const isFieldFocusRegistered = ref(false)

  const clients = ref([])
  const history = ref([])

  function setUser(payload) {
    if (payload.name) {
      userName.value = payload.name
    }
    if (payload.email) {
      userEmail.value = payload.email
    }
  }

  function fetchSampleClients() {
    axios
      .get(`data-sources/clients.json?v=3`)
      .then((result) => {
        clients.value = result?.data?.data
      })
      .catch((error) => {
        alert(error.message)
      })
  }

  function fetchSampleHistory() {
    axios
      .get(`data-sources/history.json`)
      .then((result) => {
        history.value = result?.data?.data
      })
      .catch((error) => {
        alert(error.message)
      })
  }

  return {
    userName,
    created,
    userEmail,
    stableConfig,
    userAvatar,
    isFieldFocusRegistered,
    clients,
    history,
    setUser,
    fetchSampleClients,
    fetchSampleHistory
  }
})
