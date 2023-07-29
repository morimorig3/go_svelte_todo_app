import type { RequestHandler } from './$types';

export const GET: RequestHandler = async function () {
	const res = await fetch('http://localhost:80/task/list');
	return res;
};
