new Vue({
  el: "#app",
  data: {
    gitbucketToken: "",
    gitlabToken: ""
  },
  mounted() {
    try {
      this.gitbucketToken = localStorage.getItem('gitbucketToken');
      this.gitlabToken = localStorage.getItem('gitlabToken');
    } catch (e) {
      localStorage.removeItem('gitbucketToken');
      localStorage.removeItem('gitlabToken');
    }
  },
  methods: {
    auth() {
      localStorage.setItem('gitbucketToken', this.gitbucketToken);
      localStorage.setItem('gitlabToken', this.gitlabToken);

      location.href ="/"
    }
  }
})