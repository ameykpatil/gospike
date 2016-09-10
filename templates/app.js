$(function() {

	function syntaxHighlight(json) {
	    if (typeof json != 'string') {
	         json = JSON.stringify(json, undefined, 2);
	    }
	    json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
	    return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
	        var cls = 'number';
	        if (/^"/.test(match)) {
	            if (/:$/.test(match)) {
	                cls = 'key';
	            } else {
	                cls = 'string';
	            }
	        } else if (/true|false/.test(match)) {
	            cls = 'boolean';
	        } else if (/null/.test(match)) {
	            cls = 'null';
	        }
	        return '<span class="' + cls + '">' + match + '</span>';
	    });
	}

	var recordResponse = $('#record-response');

	//***** Get record *****
	var grform = $('#get-record');
	// Set up an event listener for the contact form.
	$(grform).submit(function(e) {
		// Stop the browser from submitting the form.
		e.preventDefault();

		// Serialize the form data.
		var formData = $(grform).serialize();
		var action = $(grform).attr('action');

		var queryParams = $(grform).serializeArray().reduce(function(obj, item) {
    	obj[item.name] = item.value;
    	return obj;
		}, {});
		var key = queryParams['key']
		// Submit the form using AJAX.
		$.ajax({
			type: 'GET',
			url: action+key,
			data: formData
		})
		.done(function(response) {
			// Make sure that the formMessages div has the 'success' class.
			$(recordResponse).removeClass('error');
			$(recordResponse).addClass('success');

			// Set the message text.
			response = syntaxHighlight(response.record);
			$(recordResponse).html(response);

			// Clear the form.
			$('#gkey').val('');
		})
		.fail(function(data) {
			// Make sure that the formMessages div has the 'error' class.
			$(recordResponse).removeClass('success');
			$(recordResponse).addClass('error');

			// Set the message text.
			if (data.responseText !== '') {
				$(recordResponse).text(data.responseText);
			} else {
				$(recordResponse).text('Oops! An error occured.');
			}
		});
	});

	//***** Delete record *****
	var drform = $('#delete-record');
	
	$(drform).submit(function(e) {
		// Stop the browser from submitting the form.
		e.preventDefault();

		// Serialize the form data.
		var formData = $(drform).serialize();
		var action = $(drform).attr('action');

		var queryParams = $(drform).serializeArray().reduce(function(obj, item) {
    	obj[item.name] = item.value;
    	return obj;
		}, {});
		var key = queryParams['key']
		var namespace = queryParams['namespace']
		var set = queryParams['set']
		// Submit the form using AJAX.
		$.ajax({
			type: 'DELETE',
			url: action+key + '?' + $.param({'namespace': namespace, 'set' : set})
		})
		.done(function(response) {
			// Make sure that the formMessages div has the 'success' class.
			$(recordResponse).removeClass('error');
			$(recordResponse).addClass('success');

			// Set the message text.
			response = JSON.stringify(response);
			$(recordResponse).text(response);

			// Clear the form.
			$('#dkey').val('');
		})
		.fail(function(data) {
			// Make sure that the formMessages div has the 'error' class.
			$(recordResponse).removeClass('success');
			$(recordResponse).addClass('error');

			// Set the message text.
			if (data.responseText !== '') {
				$(recordResponse).text(data.responseText);
			} else {
				$(recordResponse).text('Oops! An error occured.');
			}
		});

	});

	// Get the form.
	var arform = $('#add-record');

	// Set up an event listener for the contact form.
	$(arform).submit(function(e) {
		// Stop the browser from submitting the form.
		e.preventDefault();

		// Serialize the form data.
		var formData = $(arform).serializeArray();
		var action = $(arform).attr('action');

		var queryParams = $(arform).serializeArray().reduce(function(obj, item) {
    	obj[item.name] = item.value;
    	return obj;
		}, {});
		var key = queryParams['key']
		var isValid = true
		var formDataObj = {};
    (function(){
        $(arform).find(":input").not("[type='submit']").not("[type='reset']").each(function(){
        	var thisInput = $(this);
        	if(thisInput.attr("name")==='record') {
        		try {
        			var recJson = JSON.parse(thisInput.val());
        			formDataObj[thisInput.attr("name")] = recJson;
        		} catch (e) {
        			isValid = false
        		}	
        	} else {
            formDataObj[thisInput.attr("name")] = thisInput.val();
          }
        });
    })();

    if(!isValid) {
    	$(recordResponse).removeClass('success');
    	$(recordResponse).addClass('error');
    	$(recordResponse).text("Can't you just enter a valid json? I might be new but I'm not dumb!");
    } else {
			// Submit the form using AJAX.
			$.ajax({
				type: 'POST',
				url: action,
				data: JSON.stringify(formDataObj),
				contentType: 'application/json',
				dataType: 'json'
			})
			.done(function(response) {
				// Make sure that the formMessages div has the 'success' class.
				$(recordResponse).removeClass('error');
				$(recordResponse).addClass('success');

				// Set the message text.
				response = syntaxHighlight(response.record);
				$(recordResponse).html(response);

				// Clear the form.
				$('#akey').val('');
				$('#arecord').val('');
			})
			.fail(function(data) {
				// Make sure that the formMessages div has the 'error' class.
				$(recordResponse).removeClass('success');
				$(recordResponse).addClass('error');

				// Set the message text.
				if (data.responseText !== '') {
					$(recordResponse).text(data.responseText);
				} else {
					$(recordResponse).text('Oops! An error occured.');
				}
			});
		}
	});

});
