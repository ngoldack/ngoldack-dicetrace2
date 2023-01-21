import { SvelteKitAuth } from "@auth/sveltekit"
import Auth0 from "@auth/core/providers/auth0"
import { env } from "$env/dynamic/private"
import { redirect } from "@sveltejs/kit";
import { sequence } from "@sveltejs/kit/hooks";

async function authorization({ event, resolve }) {
    // Protect any routes under /app
    if (event.url.pathname.startsWith('/app')) {
   const session = await event.locals.getSession();
        if (!session) {
            throw redirect(303, '/auth/signin');
        }
    }

    // If the request is still here, just proceed as normally
    const result = await resolve(event, {
        transformPageChunk: ({ html }) => html
    });
    return result;
}


export const handle = sequence(SvelteKitAuth({
    // @ts-expect-error
  providers: [Auth0({
    clientId: env.AUTH0_CLIENT_ID!,
    clientSecret: env.AUTH0_CLIENT_SECRET!,
    issuer: env.AUTH0_ISSUER
  })],
}), authorization)