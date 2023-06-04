import { Center, Text } from '@chakra-ui/react'

function NoListings () {
  return (
    <Center textAlign="center">
      <Text fontSize={{ base: '3xl', md: '4xl' }}>No Listings Found</Text>
    </Center>
  )
}

export default NoListings
