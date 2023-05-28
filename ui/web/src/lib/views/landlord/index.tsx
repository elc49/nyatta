import Head from 'next/head'

import { Container, HStack, Spacer, Show, Text } from '@chakra-ui/react'

import { usePropertyOnboarding } from '@usePropertyOnboarding'
import { Description, Location, Pricing, Units, Caretaker } from './steps'
import { Title } from './components'
import { FormSteps } from './constants'

function Landlord() {
  const { step } = usePropertyOnboarding()

  return (
    <Container>
      <Head>
        <title>Manage your properties in one place</title>
      </Head>
      <HStack my={{ base: 4, md: 6 }}>
        <Title />
        <Show above="md">
          <Spacer />
          <Text fontSize="4xl">{`${FormSteps.indexOf(step)+1}/${FormSteps.length}`}</Text>
        </Show>
      </HStack>
      {step === 'description' && <Description />}
      {step === 'location' && <Location />}
      {step === 'pricing' && <Pricing />}
      {step === 'units' && <Units />}
      {step === 'caretaker' && <Caretaker />}
    </Container>
  )
}

export default Landlord