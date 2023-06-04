import { array, object, number, string } from 'yup'
import { isValidPhoneNumber } from 'libphonenumber-js'

export const descriptionSchema = object().shape({
  name: string().trim().matches(/^[A-Za-z ]+$/i, { message: 'Property name should be alphabetic only', excludeEmptyString: true }).required('Property name is required'),
  propertyType: string().required('What is your property type?')
})

export const locationSchema = object().shape({
  town: object().shape({
    id: string(),
    label: string(),
    postalCode: string(),
    value: string()
  }).required('Town is required')
})

export const priceSchema = object().shape({
  minPrice: number().required('Minimum price is required'),
  maxPrice: number().required('Maximum price is required')
})

export const unitsSchema = object().shape({
  units: array().of(
    object().shape({
      name: string().trim().matches(/^[a-zA-Z0-9 ]+$/i, { message: 'Unit name should be alphabetic', excludeEmptyString: true }).required('Unit name required')
    })
  ).required('If you got here, your property units need to be registered')
})

export const caretakerSchema = object().shape({
  firstName: string().matches(/^[a-zA-Z ]+$/i, { message: 'First name should be alphabetic', excludeEmptyString: true }).required('First name required'),
  lastName: string().matches(/^[a-zA-Z ]+$/i, { message: 'Last name should be alphabetic', excludeEmptyString: true }).required('Last name required'),
  phoneNumber: string().matches(/^[0-9]+$/i, { message: 'Expect phone number', excludeEmptyString: true }).test('valid-phone', 'You region is not supported yet', value => {
    return isValidPhoneNumber(value, 'KE')
  }).required('Phone number required'),
  idVerification: string().required('ID verification required')
})
