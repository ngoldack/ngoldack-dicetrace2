import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async (event) => {
    console.log(await event.locals.getSession())
  return {
    session: await event.locals.getSession()
  };
}; 