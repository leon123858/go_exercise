/* eslint-disable camelcase */

exports.shorthands = undefined;

exports.up = (pgm) => {
	pgm.createTable('album', {
		id: { type: 'varchar(255)', notNull: true, primaryKey: true },
		title: { type: 'varchar(255)' },
		artist: { type: 'varchar(255)' },
		price: {
			type: 'FLOAT',
		},
	});
};

exports.down = (pgm) => {};
