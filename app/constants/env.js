'use strict';

export const IS_DEV = process.env.NODE_ENV !== 'production';
export const IS_PROD = process.env.NODE_ENV === 'production';
export const DEBUG_PROD = process.env.DEBUG_PROD === 'true';