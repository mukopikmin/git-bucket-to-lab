export const state = () => ({
  username: '',
  gitbucketUser: null,
  gitlabUser: null,
  gitbucketGroups: [],
  gitlabGroups: [],
  gitbucketUrl: '',
  gitlabUrl: '',
  gitbucketToken: '',
  gitlabToken: '',
  pairs: [],
  repo: null,
  project: null,
  error: null,
  migratable: {
    repo: false,
    issues: false,
    pulls: false
  }
})

export const mutations = {
  setGitbucketUser(state, payload) {
    state.gitbucketUser = payload.gitbucketUser
    state.username = payload.gitbucketUser.login
  },
  setGitlabUser(state, payload) {
    state.gitlabUser = payload.gitlabUser
  },
  setGitbucketGroups(state, payload) {
    state.gitbucketGroups = payload.gitbucketGroups
  },
  setGitlabGroups(state, payload) {
    state.gitlabGroups = payload.gitlabGroups
  },
  setGitbucketUrl(state, payload) {
    state.gitbucketUrl = payload.gitbucketUrl
  },
  setGitlabUrl(state, payload) {
    state.gitlabUrl = payload.gitlabUrl
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
    state.migratable.repo = payload.repoMigratable
    state.migratable.issues = payload.issuesMigratable
    state.migratable.pulls = payload.pullsMigratable
  },
  setProject(state, payload) {
    state.project = payload.project
  },
  setError(state, payload) {
    state.error = payload.error
  }
}

export const actions = {
  setGitbucketUser({ commit }, gitbucketUser) {
    commit('setGitbucketUser', { gitbucketUser })
  },
  setGitlabUser({ commit }, gitlabUser) {
    commit('setGitlabUser', { gitlabUser })
  },
  setGitbucketGroups({ commit }, gitbucketGroups) {
    commit('setGitbucketGroups', { gitbucketGroups })
  },
  setGitlabGroups({ commit }, gitlabGroups) {
    commit('setGitlabGroups', { gitlabGroups })
  },
  setGitbucketUrl({ commit }, gitbucketUrl) {
    commit('setGitbucketUrl', { gitbucketUrl })
  },
  setGitlabUrl({ commit }, gitlabUrl) {
    commit('setGitlabUrl', { gitlabUrl })
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
    commit('setRepo', { ...repo })
  },
  setProject({ commit }, project) {
    commit('setProject', { project })
  },
  setError({ commit }, error) {
    commit('setError', { error })
  }
}
