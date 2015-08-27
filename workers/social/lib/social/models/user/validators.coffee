# coffeelint: disable=indentation
# coffeelint: disable=no_implicit_braces
isEmailValid = require './emailchecker'

module.exports = class Validators

  @validateTeamDomain: (name) ->

    teamDomainPattern = ///
      ^                     # beginning of the string
      [a-z0-9]+             # one or more 0-9 and/or a-z
      (
        [-]                 # followed by a single dash
        [a-z0-9]+           # one or more (0-9 and/or a-z)
      )*                    # zero or more of the token in parentheses
      $                     # end of string
    ///

    return teamDomainPattern.test name


  createKodingError = (err) ->

    message: if 'string' is typeof err then err else err.message


  required = (field) ->

    Validators::[field] = (userData, callback) ->

      callback \
        unless userData[field]?
          createKodingError "Missed a required field: #{field}"
        else
          null

  # Since we removed first and last name requirement
  #
  # ['firstName', 'lastName'].forEach required


  agree: ({ agree }, callback) ->

    callback \
      unless agree is 'on'
        createKodingError 'You have to agree to the TOS'
      else
        null


  email: ({ email }, callback) ->

    isEmailValid email, (valid) ->

      callback \
        unless valid
          createKodingError 'Email is not valid.'
        else
          null


  username: ({ username }, callback) ->

    unless username?
      return callback createKodingError 'Missed a required field: username'

    @usernameAvailable username, (err, r) ->

      # r =
      #   forbidden    : yes/no
      #   kodingenUser : yes/no
      #   kodingUser   : yes/no

      callback \

        if err then err
        else if r.forbidden
          createKodingError 'That username is forbidden!'
        else if r.kodingUser
          createKodingError 'That username is taken!'
        else
          null


  password: ({ password, passwordConfirm }, callback) ->

    callback \

      if not password? or not passwordConfirm?
        createKodingError 'Missed a required field: password / passwordConfirm'
      else if password isnt passwordConfirm
        createKodingError 'Passwords must match'
      else if password.length < 8
        createKodingError 'Password must be at least 8 characters'
      else
        null


