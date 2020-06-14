<template>
  <div>
    <div class="row">
      <div class="col">
        <div class="form-group">
          <label>GitBucket personal access token</label>
          <input
            v-model="gitbucketTokenInput"
            type="text"
            class="form-control"
            required
          />
        </div>
      </div>
      <div class="col">
        <a
          class="btn btn-outline-primary"
          target="_blank"
          :href="`${gitbucketUrl}/root/_application`"
        >
          GitBucket personal access token
        </a>
      </div>
    </div>

    <div class="row">
      <div class="col">
        <div class="form-group">
          <label>GitLab personal access token</label>
          <input
            v-model="gitlabTokenInput"
            type="text"
            class="form-control"
            required
          />
        </div>
      </div>
      <div class="col">
        <a
          class="btn btn-outline-primary"
          target="_blank"
          :href="`${gitlabUrl}/profile/personal_access_tokens`"
        >
          GitLab personal access token
        </a>
      </div>
    </div>

    <button class="btn btn-outline-primary" @click="auth">Submit</button>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'

export default {
  data() {
    return {
      gitbucketUrl: '',
      gitlabUrl: '',
      gitbucketTokenInput: '',
      gitlabTokenInput: ''
    }
  },
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken'])
  },
  async mounted() {
    const res = await this.$axios.$get('/auth')

    this.gitbucketUrl = res.gitbucket_url
    this.gitlabUrl = res.gitlab_url
    this.gitbucketTokenInput = this.gitbucketToken
    this.gitlabTokenInput = this.gitlabToken
  },
  methods: {
    ...mapActions(['setGitbucketToken', 'setGitlabToken']),
    auth() {
      this.setGitbucketToken(this.gitbucketTokenInput)
      this.setGitlabToken(this.gitlabTokenInput)

      this.$router.push('/')
    }
  }
}
</script>

<style></style>
