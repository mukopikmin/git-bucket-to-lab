<template>
  <div>
    <ErrorMessage />

    <b-row>
      <b-col>
        <AuhorizedUser :loading="loading" />
      </b-col>
      <b-col>
        <GroupList :loading="loading" />
      </b-col>
    </b-row>

    <div class="mt-3">
      <RepoList :loading="loading" />
    </div>
  </div>
</template>

<script>
import { mapActions, mapState } from 'vuex'
import AuhorizedUser from '@/components/authorized_user'
import GroupList from '@/components/group_list'
import RepoList from '@/components/repo_list'
import ErrorMessage from '@/components/error_message'

export default {
  components: {
    AuhorizedUser,
    RepoList,
    GroupList,
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
        const pairs = await this.$axios.$get('/repos', {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        })
        this.setPairs(pairs)

        const res = await this.$axios.$get('/auth/state', {
          headers: {
            'X-GITBUCKET-TOKEN': this.gitbucketToken,
            'X-GITLAB-TOKEN': this.gitlabToken
          }
        })

        this.setGitbucketUser(res.gitbucket_user)
        this.setGitlabUser(res.gitlab_user)
        this.setGitbucketGroups(res.gitbucket_groups)
        this.setGitlabGroups(res.gitlab_groups)
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
