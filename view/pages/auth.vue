<template>
  <div>
    <div class="row">
      <div class="col">
        <div class="form-group">
          <label>GitBucket personal access token</label>
          <input
            v-model="gitbucketToken"
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
          :href="`${gitbucket_url}/root/_application`"
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
            v-model="gitlabToken"
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
          :href="`${gitlab_url}/profile/personal_access_tokens`"
        >
          GitLab personal access token
        </a>
      </div>
    </div>

    <button class="btn btn-outline-primary" @click="auth">Submit</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      gitbucketUrl: '',
      gitlabUrl: '',
      gitbucketToken: '',
      gitlabToken: ''
    }
  },
  mounted() {
    const res = axios.get('/api/auth')

    this.gitbucketUrl = res.data.gitbucket_url
    this.gitlabUrl = res.data.gitlab_url
  },
  methods: {
    auth() {
      localStorage.setItem('gitbucketToken', this.gitbucketToken)
      localStorage.setItem('gitlabToken', this.gitlabToken)

      this.$router.push('/')
    }
  }
}
</script>

<style></style>
