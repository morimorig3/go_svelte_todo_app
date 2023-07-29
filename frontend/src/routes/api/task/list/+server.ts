import { loadTaskList } from '$lib/server/task';
import { json } from '@sveltejs/kit';
import type { RequestHandler } from './$types';

export const GET: RequestHandler = async function () {
	const tasklist = await loadTaskList();
	return json(tasklist);
};
