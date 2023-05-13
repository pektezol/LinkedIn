const state = {
  user: {
    loggedIn: false,
    data: null
  }
}
const getters = {
  user (state) {
    return state.user
  }
}
const actions = {
  fetchUser () {},
  register () {},
  signInWithGoogle () {}
}
const mutations = {
  SET_LOGGED_IN (state, value) {
    state.user.loggedIn = value
  },
  SET_USER (state, data) {
    state.user.data = data
  }
}

export default {
  state,
  getters,
  actions,
  mutations
}
