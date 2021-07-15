import { shallowMount } from '@vue/test-utils'
import Home from '@/views/Home.vue'
import User from '@/views/User.vue'

describe('User.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'Hello, User'
    const wrapper = shallowMount(User, { props: { msg } })
    expect(wrapper.exists()).toBe(true)
    expect(wrapper.text()).toMatch(msg)
    expect(wrapper.find('div').exists()).toBe(true)
  })
})

describe('Home.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'Hello, World'
    const wrapper = shallowMount(Home, { props: { msg } })
    expect(wrapper.exists()).toBe(true)
    expect(wrapper.text()).toMatch(msg)
    expect(wrapper.find('div').exists()).toBe(true)
  })
})
