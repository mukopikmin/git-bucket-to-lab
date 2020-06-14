export const state = () => ({
  gitbucketToken: '',
  gitlabToken: '',
  aaa: 'aaaaa'
})

export const mutations = {
  setGitbucketToken(state, payload) {
    state.gitbucketToken = payload.token
  },
  setGitlabToken(state, payload) {
    state.gitlabToken = payload.token
  }
}

export const actions = {
  setGitbucketToken({ commit }, token) {
    commit('setGitbucketToken', { token })
  },
  setGitlabToken({ commit }, token) {
    commit('setGitlabToken', { token })
  }
}
