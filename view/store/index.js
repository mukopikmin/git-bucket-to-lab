export const state = () => ({
  gitbucketUser: '',
  gitbucketToken: '',
  gitlabToken: '',
  pairs: [],
  repo: null,
  project: null
})

export const mutations = {
  setGitbucketUser(state, payload) {
    state.gitbucketUser = payload.user
  },
  setGitbucketToken(state, payload) {
    state.gitbucketToken = payload.token
  },
  setGitlabToken(state, payload) {
    state.gitlabToken = payload.token
  },
  setPairs(state, payload) {
    state.pairs = payload.pairs
  },
  setRepo(state, payload) {
    state.repo = payload.repo
  },
  setProject(state, payload) {
    state.project = payload.project
  }
}

export const actions = {
  setGitbucketUser({ commit }, user) {
    commit('setGitbucketUser', { user })
  },
  setGitbucketToken({ commit }, token) {
    commit('setGitbucketToken', { token })
  },
  setGitlabToken({ commit }, token) {
    commit('setGitlabToken', { token })
  },
  setPairs({ commit }, pairs) {
    commit('setPairs', { pairs })
  },
  setRepo({ commit }, repo) {
    commit('setRepo', { repo })
  },
  setProject({ commit }, project) {
    commit('setProject', { project })
  }
}
