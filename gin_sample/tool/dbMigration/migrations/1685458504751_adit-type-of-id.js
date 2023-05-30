/* eslint-disable camelcase */

exports.shorthands = undefined;

exports.up = (pgm) => {
	pgm.alterColumn('album', 'id', {
		type: 'integer',
		using: 'CAST(id AS integer)',
	});
};

exports.down = (pgm) => {};
