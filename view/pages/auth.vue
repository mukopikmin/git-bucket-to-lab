<template>
  <div>
    <ErrorMessage />
    <AuthForm :loading="loading" />
  </div>
</template>

<script>
import { mapActions } from 'vuex'
import AuthForm from '@/components/auth_form'
import ErrorMessage from '@/components/error_message'

export default {
  components: {
    AuthForm,
    ErrorMessage
  },
  data() {
    return {
      loading: true
    }
  },
  async mounted() {
    try {
      const res = await this.$axios.$get('/auth')

      this.loading = false
      this.setGitbucketUrl(res.gitbucket_url)
      this.setGitlabUrl(res.gitlab_url)
      this.setError(null)
    } catch (e) {
      this.setError(e)
    }
  },
  methods: {
    ...mapActions(['setError', 'setGitbucketUrl', 'setGitlabUrl'])
  }
}
</script>

<style></style>
