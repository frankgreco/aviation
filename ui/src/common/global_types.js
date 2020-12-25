import PropTypes from 'prop-types';

// https://coolors.co/ef476f-ffd166-06d6a0-118ab2-073b4c
export const filterColorMap = {
  'n-number': '#EF476F',
  'tail number': '#EF476F',
  make: '#FFD166',
  model: '#06D6A0',
  airline: '#118AB2',
};

export const searchFilters = PropTypes.objectOf(PropTypes.shape({
  key: PropTypes.string.isRequired,
  value: PropTypes.string,
}));

export const registration = PropTypes.shape({
  n_number: PropTypes.string.isRequired,
  make: PropTypes.string.isRequired,
  Model: PropTypes.string.isRequired,
  year_manufactured: PropTypes.string.isRequired,
});
