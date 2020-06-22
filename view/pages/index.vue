<template>
  <div>
    <ErrorMessage />
    <AuhorizedUser :loading="loading" />

    <div class="mt-3">
      <RepoTable :loading="loading" />
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import AuhorizedUser from '@/components/authorized_user'
import RepoTable from '@/components/repo_table'
import ErrorMessage from '@/components/error_message'

export default {
  components: {
    AuhorizedUser,
    RepoTable,
    ErrorMessage
  },
  data() {
    return {
      gitbucketUser: null,
      gitlabUser: null,
      loading: true
    }
  },
  computed: {
    ...mapState(['gitbucketToken', 'gitlabToken']),
    isAuthoirized() {
      return !this.gitbucketToken || !this.gitlabToken
    }
  },
  mounted() {
    setTimeout(async () => {
      if (this.isAuthoirized) {
        this.$router.push('/auth')
      }

      try {
        const res = await this.$axios.$get('', {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        })

        this.setGitbucketUser(res.gitbucket_user)
        this.setGitlabUser(res.gitlab_user)
        this.setGitbucketGroups(res.gitbucket_groups)
        this.setGitlabGroups(res.gitlab_groups)
        this.setPairs(res.pairs)
        this.setError(null)
      } catch (e) {
        this.setError(e.response.data.message)
      } finally {
        this.loading = false
      }
    }, 0)
  },
  methods: {
    ...mapActions([
      'setError',
      'setPairs',
      'setGitbucketUser',
      'setGitlabUser',
      'setGitbucketGroups',
      'setGitlabGroups'
    ])
  }
}
</script>

<style></style>
