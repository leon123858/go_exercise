/* eslint-disable camelcase */

exports.shorthands = undefined;

exports.up = (pgm) => {
	pgm.alterColumn('album', 'id', {
		type: 'varchar(255)',
	});
};

exports.down = (pgm) => {};
