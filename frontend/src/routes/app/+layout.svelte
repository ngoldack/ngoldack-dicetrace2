<script>
    import auth from 'sveltekit-auth0'
	import { onMount } from "svelte/types/runtime/internal/lifecycle";
    import { getIsAuthenticatedValue, getIsAuthenticated } from 'sveltekit-auth0'

    onMount(() => {
        auth.initializeAuth0()

        if(!getIsAuthenticatedValue()) {
            auth.loginWithPopup()
        }
    })
</script>


<slot></slot>