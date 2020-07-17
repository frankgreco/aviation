import PropTypes from 'prop-types';

export const searchFilters = PropTypes.shape({
  airline: PropTypes.shape({
    enabled: PropTypes.bool.isRequired,
  }),
  make: PropTypes.shape({
    enabled: PropTypes.bool.isRequired,
  }),
  model: PropTypes.shape({
    enabled: PropTypes.bool.isRequired,
  }),
  'tail number': PropTypes.shape({
    enabled: PropTypes.bool.isRequired,
  }),
});

export const registration = PropTypes.shape({
  n_number: PropTypes.string.isRequired,
  make: PropTypes.string.isRequired,
  model: PropTypes.string.isRequired,
  year_manufactured: PropTypes.number.isRequired,
});
