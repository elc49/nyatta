import { Button, Flex, FormControl, FormErrorMessage, Input, Select } from '@chakra-ui/react'

import { useSearchListings } from '../hooks/search-listings'

function Search() {
  const { getListings, handleSubmit, register, formState: { errors } } = useSearchListings()
  const onSubmit = async (data: any) => {
    await getListings({
      variables: {
        input: {
          town: data.town,
          minPrice: Number(data.minPrice),
          maxPrice: Number(data.maxPrice),
        },
      },
    })
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <Flex
        p={5}
        gap={4}
        flexDirection={{ md: "row", base: "column" }}
      >
        <FormControl isInvalid={!!errors.town}>
          <Input
            {...register('town', {
              required: 'Town is required',
              pattern: {
                value: /^[A-Za-z ]+$/i,
                message: 'Should be a string value',
              },
            })}
            placeholder="Town"
          />
          {errors.town && <FormErrorMessage>{`${errors.town.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl isInvalid={!!errors.propertyType}>
          <Select {...register('propertyType', { required: 'Select property type' })} placeholder="Property type">
            <option value="single">Single room</option>
            <option value="studio">Studio</option>
            <option value="1">1 bedroom</option>
            <option value="2">2 bedrooms</option>
            <option value="3">3 bedrooms</option>
            <option value="4">4 bedrooms</option>
          </Select>
          {errors.propertyType && <FormErrorMessage>{`${errors.propertyType.message}`}</FormErrorMessage>}
        </FormControl>
        <FormControl>
          <Input
            {...register('minPrice')}
            type="number"
            placeholder="Min price"
            defaultValue="0"
          />
        </FormControl>
        <FormControl>
          <Input
            {...register('maxPrice')}
            type="number"
            placeholder="Max price"
          />
        </FormControl>
        <Flex>
          <Button w="100%" type="submit" colorScheme="green">
            Search
          </Button>
        </Flex>
      </Flex>
    </form>
  )
}

export default Search